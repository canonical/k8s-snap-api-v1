package apiv1

// GetClusterConfigRPC is the path for the GetClusterConfig RPC.
const GetClusterConfigRPC = "k8sd/cluster/config"

// GetClusterConfigRequest is the request message for the GetClusterConfig RPC.
type GetClusterConfigRequest struct{}

// GetClusterConfigResponse is the response message for the GetClusterConfig RPC.
type GetClusterConfigResponse struct {
	// Config is the cluster configuration.
	Config UserFacingClusterConfig `json:"status"`
	// Datastore is the datastore configuration.
	Datastore UserFacingDatastoreConfig `json:"datastore,omitempty" yaml:"datastore,omitempty"`
	// NodeTaints is a list of taints to apply to the nodes.
	// If the request comes from a control plane node, will contain the taints applied on the control plane nodes.
	// If the request comes from a worker node, will contain the taints applied on the worker nodes.
	NodeTaints *[]string `json:"node-taints,omitempty" yaml:"node-taints,omitempty"`
}
