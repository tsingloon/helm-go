// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var StorageClassKind = resource.Kind{
	Group:   "storage.k8s.io",
	Version: "v1",
	Kind:    "StorageClass",
	Scoped:  true,
}

var StorageClassResource = resource.Type{
	Kind: StorageClassKind,
	Name: "storageclasses",
}

func NewStorageClass(storageClass *storagev1.StorageClass, client resource.Client) *StorageClass {
	return &StorageClass{
		Resource: resource.NewResource(storageClass.ObjectMeta, StorageClassKind, client),
		Object:   storageClass,
	}
}

type StorageClass struct {
	*resource.Resource
	Object *storagev1.StorageClass
}

func (r *StorageClass) Delete() error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.StorageV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, StorageClassKind.Scoped).
		Resource(StorageClassResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
