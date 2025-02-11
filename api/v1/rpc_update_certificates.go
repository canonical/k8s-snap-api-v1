package apiv1

import (
	"github.com/canonical/k8s-snap-api/internal/util"
)

// UpdateCertificatesRPC is the path for the UpdateCertificates RPC.
const UpdateCertificatesRPC = "k8sd/update-certs"

// UpdateCertificatesRequest is the request message for the UpdateCertificates RPC.
type UpdateCertificatesRequest struct {
	// The client certificate to be used for the front proxy.
	FrontProxyClientCert *string `json:"front-proxy-client-crt,omitempty" yaml:"front-proxy-client-crt,omitempty"`
	// The client key to be used for the front proxy.
	FrontProxyClientKey *string `json:"front-proxy-client-key,omitempty" yaml:"front-proxy-client-key,omitempty"`
	// The admin client certificate to be used for Kubernetes services.
	AdminClientCert *string `json:"admin-client-crt,omitempty" yaml:"admin-client-crt,omitempty"`
	// The admin client key to be used for Kubernetes services.
	AdminClientKey *string `json:"admin-client-key,omitempty" yaml:"admin-client-key,omitempty"`
	// The client certificate to be used by kubelet for communicating with the kube-apiserver.
	KubeProxyClientCert *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	// The client key to be used by kubelet for communicating with the kube-apiserver.
	KubeProxyClientKey *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`
	// The client certificate to be used for the kube-scheduler.
	KubeSchedulerClientCert *string `json:"kube-scheduler-client-crt,omitempty" yaml:"kube-scheduler-client-crt,omitempty"`
	// The client key to be used for the kube-scheduler.
	KubeSchedulerClientKey *string `json:"kube-scheduler-client-key,omitempty" yaml:"kube-scheduler-client-key,omitempty"`
	// The client certificate to be used for the Kubernetes controller manager.
	KubeControllerManagerClientCert *string `json:"kube-controller-manager-client-crt,omitempty" yaml:"kube-controller-manager-client-crt,omitempty"`
	// The client key to be used for the Kubernetes controller manager.
	KubeControllerManagerClientKey *string `json:"kube-controller-manager-client-key,omitempty" yaml:"kube-controller-manager-client-key,omitempty"`

	// The certificate to be used for the kube-apiserver.
	APIServerCert *string `json:"apiserver-crt,omitempty" yaml:"apiserver-crt,omitempty"`
	// The key to be used for the kube-apiserver.
	APIServerKey *string `json:"apiserver-key,omitempty" yaml:"apiserver-key,omitempty"`
	// The certificate to be used for the kubelet.
	KubeletCert *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	// The key to be used for the kubelet.
	KubeletKey *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	// The client certificate to be used for the kubelet.
	KubeletClientCert *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	// The client key to be used for the kubelet.
	KubeletClientKey *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`
}

// UpdateCertificatesResponse is the request response for the UpdateCertificates RPC.
type UpdateCertificatesResponse struct{}

func (c *UpdateCertificatesRequest) GetFrontProxyClientCert() string {
	return util.Deref(c.FrontProxyClientCert)
}
func (c *UpdateCertificatesRequest) GetFrontProxyClientKey() string {
	return util.Deref(c.FrontProxyClientKey)
}
func (c *UpdateCertificatesRequest) GetAdminClientCert() string {
	return util.Deref(c.AdminClientCert)
}
func (c *UpdateCertificatesRequest) GetAdminClientKey() string {
	return util.Deref(c.AdminClientKey)
}
func (b *UpdateCertificatesRequest) GetKubeProxyClientCert() string {
	return util.Deref(b.KubeProxyClientCert)
}
func (b *UpdateCertificatesRequest) GetKubeProxyClientKey() string {
	return util.Deref(b.KubeProxyClientKey)
}
func (b *UpdateCertificatesRequest) GetKubeSchedulerClientCert() string {
	return util.Deref(b.KubeSchedulerClientCert)
}
func (b *UpdateCertificatesRequest) GetKubeSchedulerClientKey() string {
	return util.Deref(b.KubeSchedulerClientKey)
}
func (b *UpdateCertificatesRequest) GetKubeControllerManagerClientCert() string {
	return util.Deref(b.KubeControllerManagerClientCert)
}
func (b *UpdateCertificatesRequest) GetKubeControllerManagerClientKey() string {
	return util.Deref(b.KubeControllerManagerClientKey)
}
func (c *UpdateCertificatesRequest) GetAPIServerCert() string { return util.Deref(c.APIServerCert) }
func (c *UpdateCertificatesRequest) GetAPIServerKey() string  { return util.Deref(c.APIServerKey) }
func (c *UpdateCertificatesRequest) GetKubeletCert() string   { return util.Deref(c.KubeletCert) }
func (c *UpdateCertificatesRequest) GetKubeletKey() string    { return util.Deref(c.KubeletKey) }
func (c *UpdateCertificatesRequest) GetKubeletClientCert() string {
	return util.Deref(c.KubeletClientCert)
}
func (c *UpdateCertificatesRequest) GetKubeletClientKey() string {
	return util.Deref(c.KubeletClientKey)
}
