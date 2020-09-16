// Code generated by generate-client. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helm-client/pkg/kubernetes/resource"
)

type ReplicaSetsClient interface {
	ReplicaSets() ReplicaSetsReader
}

func NewReplicaSetsClient(resources resource.Client, filter resource.Filter) ReplicaSetsClient {
	return &replicaSetsClient{
		Client: resources,
		filter: filter,
	}
}

type replicaSetsClient struct {
	resource.Client
	filter resource.Filter
}

func (c *replicaSetsClient) ReplicaSets() ReplicaSetsReader {
	return NewReplicaSetsReader(c.Client, c.filter)
}
