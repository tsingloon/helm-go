// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type PodsReader interface {
	Get(name string) (*Pod, error)
	List() ([]*Pod, error)
}

func NewPodsReader(client resource.Client, filter resource.Filter) PodsReader {
	return &podsReader{
		Client: client,
		filter: filter,
	}
}

type podsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *podsReader) Get(name string) (*Pod, error) {
	pod := &corev1.Pod{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), PodKind.Scoped).
		Resource(PodResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(pod)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   PodKind.Group,
			Version: PodKind.Version,
			Kind:    PodKind.Kind,
		}, pod.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    PodKind.Group,
				Resource: PodResource.Name,
			}, name)
		}
	}
	return NewPod(pod, c.Client), nil
}

func (c *podsReader) List() ([]*Pod, error) {
	list := &corev1.PodList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), PodKind.Scoped).
		Resource(PodResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*Pod, 0, len(list.Items))
	for _, pod := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   PodKind.Group,
			Version: PodKind.Version,
			Kind:    PodKind.Kind,
		}, pod.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := pod
			results = append(results, NewPod(&copy, c.Client))
		}
	}
	return results, nil
}
