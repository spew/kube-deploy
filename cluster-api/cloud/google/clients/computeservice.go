package clients

import (
	compute "google.golang.org/api/compute/v1"
	"net/http"
	"net/url"
)

type ComputeService interface {
	ImagesGet(project string, image string) (*compute.Image, error)
	ImagesGetFromFamily(project string, family string) (*compute.Image, error)
	InstancesDelete(project string, zone string, targetInstance string) (*compute.Operation, error)
	InstancesGet(project string, zone string, instance string) (*compute.Instance, error)
	InstancesInsert(project string, zone string, instance *compute.Instance) (*compute.Operation, error)
	ZoneOperationsGet(project string, zone string, operation string) (*compute.Operation, error)
}

type computeServiceImpl struct {
	service *compute.Service
}

func NewComputeService(client *http.Client) (ComputeService, error) {
	return newComputeServiceImpl(client)
}

func NewComputeServiceForURL(client *http.Client, baseURL string) (ComputeService, error) {
	computeServiceImpl, err := newComputeServiceImpl(client)
	if err != nil {
		return nil, err
	}
	url, err := url.Parse(computeServiceImpl.service.BasePath)
	if err != nil {
		return nil, err
	}
	computeServiceImpl.service.BasePath = baseURL + url.Path
	return computeServiceImpl, err
}

func newComputeServiceImpl(client *http.Client) (*computeServiceImpl, error) {
	service, err := compute.New(client)
	if err != nil {
		return nil, err
	}
	return &computeServiceImpl{
		service: service,
	}, nil
}

func (c *computeServiceImpl) ImagesGet(project string, image string) (*compute.Image, error) {
	return c.service.Images.Get(project, image).Do()
}

func (c *computeServiceImpl) ImagesGetFromFamily(project string, family string) (*compute.Image, error) {
	return c.service.Images.GetFromFamily(project, family).Do()
}

func (c *computeServiceImpl) InstancesDelete(project string, zone string, targetInstance string) (*compute.Operation, error) {
	return c.service.Instances.Delete(project, zone, targetInstance).Do()
}

func (c *computeServiceImpl) InstancesGet(project string, zone string, instance string) (*compute.Instance, error) {
	return c.service.Instances.Get(project, zone, instance).Do()
}

func (c *computeServiceImpl) InstancesInsert(project string, zone string, instance *compute.Instance) (*compute.Operation, error) {
	return c.service.Instances.Insert(project, zone, instance).Do()
}

func (c *computeServiceImpl) ZoneOperationsGet(project string, zone string, operation string) (*compute.Operation, error) {
	return c.service.ZoneOperations.Get(project, zone, operation).Do()
}
