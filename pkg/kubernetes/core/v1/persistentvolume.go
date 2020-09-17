// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-go/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var PersistentVolumeKind = resource.Kind{
	Group:   "",
	Version: "v1",
	Kind:    "PersistentVolume",
	Scoped:  true,
}

var PersistentVolumeResource = resource.Type{
	Kind: PersistentVolumeKind,
	Name: "persistentvolumes",
}

func NewPersistentVolume(persistentVolume *corev1.PersistentVolume, client resource.Client) *PersistentVolume {
	return &PersistentVolume{
		Resource: resource.NewResource(persistentVolume.ObjectMeta, PersistentVolumeKind, client),
		Object:   persistentVolume,
	}
}

type PersistentVolume struct {
	*resource.Resource
	Object *corev1.PersistentVolume
}

func (r *PersistentVolume) Delete() error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.CoreV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, PersistentVolumeKind.Scoped).
		Resource(PersistentVolumeResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
