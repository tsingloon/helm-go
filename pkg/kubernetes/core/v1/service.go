// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-go/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var ServiceKind = resource.Kind{
	Group:   "",
	Version: "v1",
	Kind:    "Service",
	Scoped:  true,
}

var ServiceResource = resource.Type{
	Kind: ServiceKind,
	Name: "services",
}

func NewService(service *corev1.Service, client resource.Client) *Service {
	return &Service{
		Resource:           resource.NewResource(service.ObjectMeta, ServiceKind, client),
		Object:             service,
		EndpointsReference: NewEndpointsReference(client, resource.NewUIDFilter(service.UID)),
	}
}

type Service struct {
	*resource.Resource
	Object *corev1.Service
	EndpointsReference
}

func (r *Service) Delete() error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.CoreV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, ServiceKind.Scoped).
		Resource(ServiceResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
