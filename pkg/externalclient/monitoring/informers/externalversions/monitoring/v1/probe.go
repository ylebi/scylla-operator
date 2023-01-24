// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	monitoringv1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	versioned "github.com/scylladb/scylla-operator/pkg/externalclient/monitoring/clientset/versioned"
	internalinterfaces "github.com/scylladb/scylla-operator/pkg/externalclient/monitoring/informers/externalversions/internalinterfaces"
	v1 "github.com/scylladb/scylla-operator/pkg/externalclient/monitoring/listers/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProbeInformer provides access to a shared informer and lister for
// Probes.
type ProbeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProbeLister
}

type probeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProbeInformer constructs a new informer for Probe type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProbeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProbeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProbeInformer constructs a new informer for Probe type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProbeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().Probes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().Probes(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1.Probe{},
		resyncPeriod,
		indexers,
	)
}

func (f *probeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProbeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *probeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1.Probe{}, f.defaultInformer)
}

func (f *probeInformer) Lister() v1.ProbeLister {
	return v1.NewProbeLister(f.Informer().GetIndexer())
}
