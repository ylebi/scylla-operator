// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScyllaDBClusterLister helps list ScyllaDBClusters.
// All objects returned here must be treated as read-only.
type ScyllaDBClusterLister interface {
	// List lists all ScyllaDBClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScyllaDBCluster, err error)
	// ScyllaDBClusters returns an object that can list and get ScyllaDBClusters.
	ScyllaDBClusters(namespace string) ScyllaDBClusterNamespaceLister
	ScyllaDBClusterListerExpansion
}

// scyllaDBClusterLister implements the ScyllaDBClusterLister interface.
type scyllaDBClusterLister struct {
	indexer cache.Indexer
}

// NewScyllaDBClusterLister returns a new ScyllaDBClusterLister.
func NewScyllaDBClusterLister(indexer cache.Indexer) ScyllaDBClusterLister {
	return &scyllaDBClusterLister{indexer: indexer}
}

// List lists all ScyllaDBClusters in the indexer.
func (s *scyllaDBClusterLister) List(selector labels.Selector) (ret []*v1alpha1.ScyllaDBCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScyllaDBCluster))
	})
	return ret, err
}

// ScyllaDBClusters returns an object that can list and get ScyllaDBClusters.
func (s *scyllaDBClusterLister) ScyllaDBClusters(namespace string) ScyllaDBClusterNamespaceLister {
	return scyllaDBClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScyllaDBClusterNamespaceLister helps list and get ScyllaDBClusters.
// All objects returned here must be treated as read-only.
type ScyllaDBClusterNamespaceLister interface {
	// List lists all ScyllaDBClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScyllaDBCluster, err error)
	// Get retrieves the ScyllaDBCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScyllaDBCluster, error)
	ScyllaDBClusterNamespaceListerExpansion
}

// scyllaDBClusterNamespaceLister implements the ScyllaDBClusterNamespaceLister
// interface.
type scyllaDBClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScyllaDBClusters in the indexer for a given namespace.
func (s scyllaDBClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScyllaDBCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScyllaDBCluster))
	})
	return ret, err
}

// Get retrieves the ScyllaDBCluster from the indexer for a given namespace and name.
func (s scyllaDBClusterNamespaceLister) Get(name string) (*v1alpha1.ScyllaDBCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scylladbcluster"), name)
	}
	return obj.(*v1alpha1.ScyllaDBCluster), nil
}