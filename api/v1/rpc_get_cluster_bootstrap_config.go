package apiv1

// GetClusterBootstrapConfigRPC is the path for the GetClusterBootstrapConfig RPC.
const GetClusterBootstrapConfigRPC = "k8sd/cluster/config/bootstrap"

// GetClusterBootstrapConfigRequest is the request message for the GetClusterBootstrapConfig RPC.
type GetClusterBootstrapConfigRequest struct{}

// GetClusterBootstrapConfigResponse is the response message for the GetClusterBootstrapConfig RPC.
type GetClusterBootstrapConfigResponse struct {
	// Config is the cluster configuration.
	ClusterConfig UserFacingClusterConfig `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`
	// Datastore is the datastore configuration.
	Datastore UserFacingDatastoreConfig `json:"datastore,omitempty" yaml:"datastore,omitempty"`
	// NodeTaints is a list of taints to apply to the nodes.
	// If the request comes from a control plane node, will contain the taints applied on the control plane nodes.
	// If the request comes from a worker node, will contain the taints applied on the worker nodes.
	NodeTaints *[]string `json:"nodeTaints,omitempty" yaml:"nodeTaints,omitempty"`
}
