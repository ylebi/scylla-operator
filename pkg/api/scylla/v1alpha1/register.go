package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupName     = "scylla.scylladb.com"
	GroupVersion  = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// Install is a function which adds this version to a scheme
	Install = schemeBuilder.AddToScheme

	// SchemeGroupVersion generated code relies on this name
	// Deprecated
	SchemeGroupVersion = GroupVersion
	// AddToScheme exists solely to keep the old generators creating valid code
	// DEPRECATED
	AddToScheme = schemeBuilder.AddToScheme
)

var (
	ScyllaDBDatacenterGVK = GroupVersion.WithKind("ScyllaDBDatacenter")
	ScyllaDBClusterGVK    = GroupVersion.WithKind("ScyllaDBCluster")
)

// Resource generated code relies on this being here, but it logically belongs to the group
// DEPRECATED
func Resource(resource string) schema.GroupResource {
	return schema.GroupResource{Group: GroupName, Resource: resource}
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&NodeConfig{},
		&NodeConfigList{},
		&ScyllaOperatorConfig{},
		&ScyllaOperatorConfigList{},
		&ScyllaDBMonitoring{},
		&ScyllaDBMonitoringList{},
		&ScyllaDBDatacenter{},
		&ScyllaDBDatacenterList{},
		&ScyllaDBCluster{},
		&ScyllaDBClusterList{},
		&RemoteKubernetesCluster{},
		&RemoteKubernetesClusterList{},
		&RemoteOwner{},
		&RemoteOwnerList{},
		&ScyllaDBManagerClusterRegistration{},
		&ScyllaDBManagerClusterRegistrationList{},
		&ScyllaDBManagerTask{},
		&ScyllaDBManagerTaskList{},
	)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
