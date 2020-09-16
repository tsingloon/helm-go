// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type DaemonSetsReader interface {
	Get(name string) (*DaemonSet, error)
	List() ([]*DaemonSet, error)
}

func NewDaemonSetsReader(client resource.Client, filter resource.Filter) DaemonSetsReader {
	return &daemonSetsReader{
		Client: client,
		filter: filter,
	}
}

type daemonSetsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *daemonSetsReader) Get(name string) (*DaemonSet, error) {
	daemonSet := &appsv1.DaemonSet{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.AppsV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), DaemonSetKind.Scoped).
		Resource(DaemonSetResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(daemonSet)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   DaemonSetKind.Group,
			Version: DaemonSetKind.Version,
			Kind:    DaemonSetKind.Kind,
		}, daemonSet.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    DaemonSetKind.Group,
				Resource: DaemonSetResource.Name,
			}, name)
		}
	}
	return NewDaemonSet(daemonSet, c.Client), nil
}

func (c *daemonSetsReader) List() ([]*DaemonSet, error) {
	list := &appsv1.DaemonSetList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.AppsV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), DaemonSetKind.Scoped).
		Resource(DaemonSetResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*DaemonSet, 0, len(list.Items))
	for _, daemonSet := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   DaemonSetKind.Group,
			Version: DaemonSetKind.Version,
			Kind:    DaemonSetKind.Kind,
		}, daemonSet.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := daemonSet
			results = append(results, NewDaemonSet(&copy, c.Client))
		}
	}
	return results, nil
}
