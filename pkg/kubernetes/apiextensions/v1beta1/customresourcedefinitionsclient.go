// Code generated by generate-client. DO NOT EDIT.

package v1beta1

import (
	"github.com/onosproject/helm-go/pkg/kubernetes/resource"
)

type CustomResourceDefinitionsClient interface {
	CustomResourceDefinitions() CustomResourceDefinitionsReader
}

func NewCustomResourceDefinitionsClient(resources resource.Client, filter resource.Filter) CustomResourceDefinitionsClient {
	return &customResourceDefinitionsClient{
		Client: resources,
		filter: filter,
	}
}

type customResourceDefinitionsClient struct {
	resource.Client
	filter resource.Filter
}

func (c *customResourceDefinitionsClient) CustomResourceDefinitions() CustomResourceDefinitionsReader {
	return NewCustomResourceDefinitionsReader(c.Client, c.filter)
}
