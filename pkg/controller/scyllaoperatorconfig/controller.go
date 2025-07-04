package scyllaoperatorconfig

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	scyllav1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	scyllav1alpha1client "github.com/scylladb/scylla-operator/pkg/client/scylla/clientset/versioned/typed/scylla/v1alpha1"
	scyllav1alpha1informers "github.com/scylladb/scylla-operator/pkg/client/scylla/informers/externalversions/scylla/v1alpha1"
	scyllav1alpha1listers "github.com/scylladb/scylla-operator/pkg/client/scylla/listers/scylla/v1alpha1"
	"github.com/scylladb/scylla-operator/pkg/controllerhelpers"
	"github.com/scylladb/scylla-operator/pkg/kubeinterfaces"
	"github.com/scylladb/scylla-operator/pkg/naming"
	"github.com/scylladb/scylla-operator/pkg/scheme"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	apimachineryutilerrors "k8s.io/apimachinery/pkg/util/errors"
	apimachineryutilruntime "k8s.io/apimachinery/pkg/util/runtime"
	apimachineryutilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const (
	ControllerName = "ScyllaOperatorConfigController"
	// maxSyncDuration enforces preemption. Do not raise the value! Controllers shouldn't actively wait,
	// but rather use the queue.
	maxSyncDuration           = 30 * time.Second
	clusterDomainPollInterval = 5 * time.Minute
)

var (
	keyFunc                           = cache.DeletionHandlingMetaNamespaceKeyFunc
	scyllaOperatorConfigControllerGVK = scyllav1.GroupVersion.WithKind("ScyllaOperatorConfig")
)

type GetClusterDomainFunc func(context.Context) (string, error)

type Controller struct {
	kubeClient   kubernetes.Interface
	scyllaClient scyllav1alpha1client.ScyllaV1alpha1Interface

	scyllaOperatorConfigLister scyllav1alpha1listers.ScyllaOperatorConfigLister

	cachesToSync []cache.InformerSynced

	getClusterDomainFunc GetClusterDomainFunc

	eventRecorder record.EventRecorder

	queue    workqueue.RateLimitingInterface
	handlers *controllerhelpers.Handlers[*scyllav1alpha1.ScyllaOperatorConfig]

	wg sync.WaitGroup
}

func NewController(
	kubeClient kubernetes.Interface,
	scyllaClient scyllav1alpha1client.ScyllaV1alpha1Interface,
	scyllaOperatorConfigInformer scyllav1alpha1informers.ScyllaOperatorConfigInformer,
	getClusterDomain GetClusterDomainFunc,
) (*Controller, error) {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&corev1client.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})

	opc := &Controller{
		kubeClient:   kubeClient,
		scyllaClient: scyllaClient,

		scyllaOperatorConfigLister: scyllaOperatorConfigInformer.Lister(),

		cachesToSync: []cache.InformerSynced{
			scyllaOperatorConfigInformer.Informer().HasSynced,
		},

		getClusterDomainFunc: getClusterDomain,

		eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "orphanedpv-controller"}),

		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "orphanedpv"),
	}

	var err error
	opc.handlers, err = controllerhelpers.NewHandlers[*scyllav1alpha1.ScyllaOperatorConfig](
		opc.queue,
		keyFunc,
		scheme.Scheme,
		scyllaOperatorConfigControllerGVK,
		kubeinterfaces.GlobalGetList[*scyllav1alpha1.ScyllaOperatorConfig]{
			GetFunc: func(name string) (*scyllav1alpha1.ScyllaOperatorConfig, error) {
				return opc.scyllaOperatorConfigLister.Get(name)
			},
			ListFunc: func(selector labels.Selector) (ret []*scyllav1alpha1.ScyllaOperatorConfig, err error) {
				return opc.scyllaOperatorConfigLister.List(selector)
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("can't create handlers: %w", err)
	}

	scyllaOperatorConfigInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    opc.addScyllaOperatorConfig,
		UpdateFunc: opc.updateScyllaOperatorConfig,
		DeleteFunc: opc.deleteScyllaOperatorConfig,
	})

	opc.queue.Add(naming.SingletonName)

	return opc, nil
}

func (opc *Controller) processNextItem(ctx context.Context) bool {
	key, quit := opc.queue.Get()
	if quit {
		return false
	}
	defer opc.queue.Done(key)

	ctx, cancel := context.WithTimeout(ctx, maxSyncDuration)
	defer cancel()
	syncErr := opc.sync(ctx)
	if syncErr == nil {
		opc.queue.Forget(key)
		return true
	}

	// Make sure we always have an aggregate to process and all nested errors are flattened.
	allErrors := apimachineryutilerrors.Flatten(apimachineryutilerrors.NewAggregate([]error{syncErr}))
	var remainingErrors []error
	for _, err := range allErrors.Errors() {
		switch {
		case errors.Is(err, &controllerhelpers.RequeueError{}):
			klog.V(2).InfoS("Re-queuing for recheck", "Key", key, "Reason", err)

		case apierrors.IsConflict(err):
			klog.V(2).InfoS("Hit conflict, will retry in a bit", "Key", key, "Error", err)

		case apierrors.IsAlreadyExists(err):
			klog.V(2).InfoS("Hit already exists, will retry in a bit", "Key", key, "Error", err)

		default:
			remainingErrors = append(remainingErrors, err)
		}
	}

	err := apimachineryutilerrors.NewAggregate(remainingErrors)
	if err != nil {
		apimachineryutilruntime.HandleError(fmt.Errorf("syncing key '%v' failed: %v", key, err))
	}

	opc.queue.AddRateLimited(key)

	return true
}

func (opc *Controller) runWorker(ctx context.Context) {
	for opc.processNextItem(ctx) {
	}
}

func (opc *Controller) Run(ctx context.Context, workers int) {
	defer apimachineryutilruntime.HandleCrash()

	klog.InfoS("Starting controller", "controller", ControllerName)

	defer func() {
		klog.InfoS("Shutting down controller", "controller", ControllerName)
		opc.queue.ShutDown()
		opc.wg.Wait()
		klog.InfoS("Shut down controller", "controller", ControllerName)
	}()

	if !cache.WaitForNamedCacheSync(ControllerName, ctx.Done(), opc.cachesToSync...) {
		return
	}

	for range workers {
		opc.wg.Add(1)
		go func() {
			defer opc.wg.Done()
			apimachineryutilwait.UntilWithContext(ctx, opc.runWorker, time.Second)
		}()
	}

	opc.wg.Add(1)
	go func() {
		defer opc.wg.Done()
		apimachineryutilwait.UntilWithContext(ctx, func(ctx context.Context) {
			klog.V(4).InfoS("Periodically enqueuing %q ScyllaOperatorConfig", naming.SingletonName)

			key, err := keyFunc(&metav1.ObjectMeta{
				Namespace: "",
				Name:      naming.SingletonName,
			})
			if err != nil {
				apimachineryutilruntime.HandleError(fmt.Errorf("couldn't get key for object name %#v: %v", naming.SingletonName, err))
				return
			}

			opc.queue.Add(key)
		}, clusterDomainPollInterval)
	}()

	<-ctx.Done()
}

func (opc *Controller) addScyllaOperatorConfig(obj interface{}) {
	opc.handlers.HandleAdd(
		obj.(*scyllav1alpha1.ScyllaOperatorConfig),
		opc.handlers.Enqueue,
	)
}

func (opc *Controller) updateScyllaOperatorConfig(old, cur interface{}) {
	opc.handlers.HandleUpdate(
		old.(*scyllav1alpha1.ScyllaOperatorConfig),
		cur.(*scyllav1alpha1.ScyllaOperatorConfig),
		opc.handlers.Enqueue,
		opc.deleteScyllaOperatorConfig,
	)
}

func (opc *Controller) deleteScyllaOperatorConfig(obj interface{}) {
	opc.handlers.HandleDelete(
		obj,
		opc.handlers.Enqueue,
	)
}
