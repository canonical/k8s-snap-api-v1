package apiv1

import (
	"github.com/canonical/k8s-snap-api/internal/util"
)

// UpdateCertificatesRPC is the path for the UpdateCertificates RPC.
const UpdateCertificatesRPC = "k8sd/refresh-certs/update"

// UpdateCertificatesRequest is the request message for the UpdateCertificates RPC.
type UpdateCertificatesRequest struct {
	FrontProxyClientCert *string `json:"front-proxy-client-crt,omitempty" yaml:"front-proxy-client-crt,omitempty"`
	// The client key to be used for the front proxy.
	FrontProxyClientKey *string `json:"front-proxy-client-key,omitempty" yaml:"front-proxy-client-key,omitempty"`
	// The client certificate to be used by kubelet for communicating with the kube-apiserver.
	APIServerKubeletClientCert *string `json:"apiserver-kubelet-client-crt,omitempty" yaml:"apiserver-kubelet-client-crt,omitempty"`
	// The client key to be used by kubelet for communicating with the kube-apiserver.
	APIServerKubeletClientKey *string `json:"apiserver-kubelet-client-key,omitempty" yaml:"apiserver-kubelet-client-key,omitempty"`
	// The admin client certificate to be used for Kubernetes services.
	AdminClientCert *string `json:"admin-client-crt,omitempty" yaml:"admin-client-crt,omitempty"`
	// The admin client key to be used for Kubernetes services.
	AdminClientKey *string `json:"admin-client-key,omitempty" yaml:"admin-client-key,omitempty"`
	// The client certificate to be used for the kube-proxy.
	KubeProxyClientCert *string `json:"kube-proxy-client-crt,omitempty" yaml:"kube-proxy-client-crt,omitempty"`
	// The client key to be used for the kube-proxy.
	KubeProxyClientKey *string `json:"kube-proxy-client-key,omitempty" yaml:"kube-proxy-client-key,omitempty"`
	// The client certificate to be used for the kube-scheduler.
	KubeSchedulerClientCert *string `json:"kube-scheduler-client-crt,omitempty" yaml:"kube-scheduler-client-crt,omitempty"`
	// The client key to be used for the kube-scheduler.
	KubeSchedulerClientKey *string `json:"kube-scheduler-client-key,omitempty" yaml:"kube-scheduler-client-key,omitempty"`
	// The client certificate to be used for the Kubernetes controller manager.
	KubeControllerManagerClientCert *string `json:"kube-controller-manager-client-crt,omitempty" yaml:"kube-controller-manager-client-crt,omitempty"`
	// The client key to be used for the Kubernetes controller manager.
	KubeControllerManagerClientKey *string `json:"kube-controller-manager-client-key,omitempty" yaml:"kube-ControllerManager-client-key,omitempty"`
	// The certificate to be used for the kube-apiserver.
	APIServerCert *string `json:"apiserver-crt,omitempty" yaml:"apiserver-crt,omitempty"`
	// The key to be used for the kube-apiserver.
	APIServerKey *string `json:"apiserver-key,omitempty" yaml:"apiserver-key,omitempty"`
	// The certificate to be used for the kubelet.
	KubeletCert *string `json:"kubelet-crt,omitempty" yaml:"kubelet-crt,omitempty"`
	// The key to be used for the kubelet.
	KubeletKey *string `json:"kubelet-key,omitempty" yaml:"kubelet-key,omitempty"`
	// The certificate to be used for the kubelet client.
	KubeletClientCert *string `json:"kubelet-client-crt,omitempty" yaml:"kubelet-client-crt,omitempty"`
	// The key to be used for the kubelet client.
	KubeletClientKey *string `json:"kubelet-client-key,omitempty" yaml:"kubelet-client-key,omitempty"`
}

// UpdateCertificatesResponse is the request response for the UpdateCertificates RPC.
type UpdateCertificatesResponse struct{}

func (u *UpdateCertificatesRequest) GetFrontProxyClientCert() string {
	return util.Deref(u.FrontProxyClientCert)
}
func (u *UpdateCertificatesRequest) GetFrontProxyClientKey() string {
	return util.Deref(u.FrontProxyClientKey)
}
func (u *UpdateCertificatesRequest) GetAPIServerKubeletClientCert() string {
	return util.Deref(u.APIServerKubeletClientCert)
}
func (u *UpdateCertificatesRequest) GetAPIServerKubeletClientKey() string {
	return util.Deref(u.APIServerKubeletClientKey)
}
func (u *UpdateCertificatesRequest) GetAdminClientCert() string {
	return util.Deref(u.AdminClientCert)
}
func (u *UpdateCertificatesRequest) GetAdminClientKey() string {
	return util.Deref(u.AdminClientKey)
}
func (u *UpdateCertificatesRequest) GetKubeProxyClientCert() string {
	return util.Deref(u.KubeProxyClientCert)
}
func (u *UpdateCertificatesRequest) GetKubeProxyClientKey() string {
	return util.Deref(u.KubeProxyClientKey)
}
func (u *UpdateCertificatesRequest) GetKubeSchedulerClientCert() string {
	return util.Deref(u.KubeSchedulerClientCert)
}
func (u *UpdateCertificatesRequest) GetKubeSchedulerClientKey() string {
	return util.Deref(u.KubeSchedulerClientKey)
}
func (u *UpdateCertificatesRequest) GetKubeControllerManagerClientCert() string {
	return util.Deref(u.KubeControllerManagerClientCert)
}
func (u *UpdateCertificatesRequest) GetKubeControllerManagerClientKey() string {
	return util.Deref(u.KubeControllerManagerClientKey)
}
func (u *UpdateCertificatesRequest) GetAPIServerCert() string { return util.Deref(u.APIServerCert) }
func (u *UpdateCertificatesRequest) GetAPIServerKey() string  { return util.Deref(u.APIServerKey) }
func (u *UpdateCertificatesRequest) GetKubeletCert() string   { return util.Deref(u.KubeletCert) }
func (u *UpdateCertificatesRequest) GetKubeletKey() string    { return util.Deref(u.KubeletKey) }
func (u *UpdateCertificatesRequest) GetKubeletClientCert() string {
	return util.Deref(u.KubeletClientCert)
}
func (u *UpdateCertificatesRequest) GetKubeletClientKey() string {
	return util.Deref(u.KubeletClientKey)
}
