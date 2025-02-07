// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	scyllav1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	scheme "github.com/scylladb/scylla-operator/pkg/client/scylla/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ScyllaOperatorConfigsGetter has a method to return a ScyllaOperatorConfigInterface.
// A group's client should implement this interface.
type ScyllaOperatorConfigsGetter interface {
	ScyllaOperatorConfigs() ScyllaOperatorConfigInterface
}

// ScyllaOperatorConfigInterface has methods to work with ScyllaOperatorConfig resources.
type ScyllaOperatorConfigInterface interface {
	Create(ctx context.Context, scyllaOperatorConfig *scyllav1alpha1.ScyllaOperatorConfig, opts v1.CreateOptions) (*scyllav1alpha1.ScyllaOperatorConfig, error)
	Update(ctx context.Context, scyllaOperatorConfig *scyllav1alpha1.ScyllaOperatorConfig, opts v1.UpdateOptions) (*scyllav1alpha1.ScyllaOperatorConfig, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, scyllaOperatorConfig *scyllav1alpha1.ScyllaOperatorConfig, opts v1.UpdateOptions) (*scyllav1alpha1.ScyllaOperatorConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*scyllav1alpha1.ScyllaOperatorConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*scyllav1alpha1.ScyllaOperatorConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *scyllav1alpha1.ScyllaOperatorConfig, err error)
	ScyllaOperatorConfigExpansion
}

// scyllaOperatorConfigs implements ScyllaOperatorConfigInterface
type scyllaOperatorConfigs struct {
	*gentype.ClientWithList[*scyllav1alpha1.ScyllaOperatorConfig, *scyllav1alpha1.ScyllaOperatorConfigList]
}

// newScyllaOperatorConfigs returns a ScyllaOperatorConfigs
func newScyllaOperatorConfigs(c *ScyllaV1alpha1Client) *scyllaOperatorConfigs {
	return &scyllaOperatorConfigs{
		gentype.NewClientWithList[*scyllav1alpha1.ScyllaOperatorConfig, *scyllav1alpha1.ScyllaOperatorConfigList](
			"scyllaoperatorconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *scyllav1alpha1.ScyllaOperatorConfig { return &scyllav1alpha1.ScyllaOperatorConfig{} },
			func() *scyllav1alpha1.ScyllaOperatorConfigList { return &scyllav1alpha1.ScyllaOperatorConfigList{} },
		),
	}
}
