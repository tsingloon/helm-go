// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-go/pkg/kubernetes/resource"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type RolesReader interface {
	Get(name string) (*Role, error)
	List() ([]*Role, error)
}

func NewRolesReader(client resource.Client, filter resource.Filter) RolesReader {
	return &rolesReader{
		Client: client,
		filter: filter,
	}
}

type rolesReader struct {
	resource.Client
	filter resource.Filter
}

func (c *rolesReader) Get(name string) (*Role, error) {
	role := &rbacv1.Role{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.RbacV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), RoleKind.Scoped).
		Resource(RoleResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(role)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   RoleKind.Group,
			Version: RoleKind.Version,
			Kind:    RoleKind.Kind,
		}, role.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    RoleKind.Group,
				Resource: RoleResource.Name,
			}, name)
		}
	}
	return NewRole(role, c.Client), nil
}

func (c *rolesReader) List() ([]*Role, error) {
	list := &rbacv1.RoleList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.RbacV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), RoleKind.Scoped).
		Resource(RoleResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*Role, 0, len(list.Items))
	for _, role := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   RoleKind.Group,
			Version: RoleKind.Version,
			Kind:    RoleKind.Kind,
		}, role.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := role
			results = append(results, NewRole(&copy, c.Client))
		}
	}
	return results, nil
}
