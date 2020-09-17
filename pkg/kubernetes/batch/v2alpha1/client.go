// Code generated by generate-client. DO NOT EDIT.

package v2alpha1

import (
	"github.com/onosproject/helm-go/pkg/kubernetes/resource"
)

type Client interface {
	CronJobsClient
}

func NewClient(resources resource.Client, filter resource.Filter) Client {
	return &client{
		Client:         resources,
		CronJobsClient: NewCronJobsClient(resources, filter),
	}
}

type client struct {
	resource.Client
	CronJobsClient
}
