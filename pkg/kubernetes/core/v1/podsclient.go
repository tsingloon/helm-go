// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
)

type PodsClient interface {
	Pods() PodsReader
}

func NewPodsClient(resources resource.Client, filter resource.Filter) PodsClient {
	return &podsClient{
		Client: resources,
		filter: filter,
	}
}

type podsClient struct {
	resource.Client
	filter resource.Filter
}

func (c *podsClient) Pods() PodsReader {
	return NewPodsReader(c.Client, c.filter)
}
