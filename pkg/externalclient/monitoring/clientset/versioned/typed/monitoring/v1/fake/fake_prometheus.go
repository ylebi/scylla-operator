// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"

	v1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	monitoringv1 "github.com/scylladb/scylla-operator/pkg/externalclient/monitoring/clientset/versioned/typed/monitoring/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
	testing "k8s.io/client-go/testing"
)

// fakePrometheuses implements PrometheusInterface
type fakePrometheuses struct {
	*gentype.FakeClientWithList[*v1.Prometheus, *v1.PrometheusList]
	Fake *FakeMonitoringV1
}

func newFakePrometheuses(fake *FakeMonitoringV1, namespace string) monitoringv1.PrometheusInterface {
	return &fakePrometheuses{
		gentype.NewFakeClientWithList[*v1.Prometheus, *v1.PrometheusList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("prometheuses"),
			v1.SchemeGroupVersion.WithKind("Prometheus"),
			func() *v1.Prometheus { return &v1.Prometheus{} },
			func() *v1.PrometheusList { return &v1.PrometheusList{} },
			func(dst, src *v1.PrometheusList) { dst.ListMeta = src.ListMeta },
			func(list *v1.PrometheusList) []*v1.Prometheus { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.PrometheusList, items []*v1.Prometheus) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

// GetScale takes name of the prometheus, and returns the corresponding scale object, and an error if there is any.
func (c *fakePrometheuses) GetScale(ctx context.Context, prometheusName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(c.Resource(), c.Namespace(), "scale", prometheusName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *fakePrometheuses) UpdateScale(ctx context.Context, prometheusName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(c.Resource(), "scale", c.Namespace(), scale, opts), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}
