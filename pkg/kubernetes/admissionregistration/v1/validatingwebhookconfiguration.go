// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var ValidatingWebhookConfigurationKind = resource.Kind{
	Group:   "admissionregistration.k8s.io",
	Version: "v1",
	Kind:    "ValidatingWebhookConfiguration",
	Scoped:  true,
}

var ValidatingWebhookConfigurationResource = resource.Type{
	Kind: ValidatingWebhookConfigurationKind,
	Name: "validatingwebhookconfigurations",
}

func NewValidatingWebhookConfiguration(validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration, client resource.Client) *ValidatingWebhookConfiguration {
	return &ValidatingWebhookConfiguration{
		Resource: resource.NewResource(validatingWebhookConfiguration.ObjectMeta, ValidatingWebhookConfigurationKind, client),
		Object:   validatingWebhookConfiguration,
	}
}

type ValidatingWebhookConfiguration struct {
	*resource.Resource
	Object *admissionregistrationv1.ValidatingWebhookConfiguration
}

func (r *ValidatingWebhookConfiguration) Delete() error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.AdmissionregistrationV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, ValidatingWebhookConfigurationKind.Scoped).
		Resource(ValidatingWebhookConfigurationResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
