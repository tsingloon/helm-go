// Code generated by generate-client. DO NOT EDIT.

package v1beta1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type IngressesReader interface {
	Get(name string) (*Ingress, error)
	List() ([]*Ingress, error)
}

func NewIngressesReader(client resource.Client, filter resource.Filter) IngressesReader {
	return &ingressesReader{
		Client: client,
		filter: filter,
	}
}

type ingressesReader struct {
	resource.Client
	filter resource.Filter
}

func (c *ingressesReader) Get(name string) (*Ingress, error) {
	ingress := &networkingv1beta1.Ingress{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.NetworkingV1beta1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), IngressKind.Scoped).
		Resource(IngressResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(ingress)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   IngressKind.Group,
			Version: IngressKind.Version,
			Kind:    IngressKind.Kind,
		}, ingress.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    IngressKind.Group,
				Resource: IngressResource.Name,
			}, name)
		}
	}
	return NewIngress(ingress, c.Client), nil
}

func (c *ingressesReader) List() ([]*Ingress, error) {
	list := &networkingv1beta1.IngressList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.NetworkingV1beta1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), IngressKind.Scoped).
		Resource(IngressResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*Ingress, 0, len(list.Items))
	for _, ingress := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   IngressKind.Group,
			Version: IngressKind.Version,
			Kind:    IngressKind.Kind,
		}, ingress.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := ingress
			results = append(results, NewIngress(&copy, c.Client))
		}
	}
	return results, nil
}
