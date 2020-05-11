package application

import (
	"fmt"
	"github.com/reuben-baek/clean-go-application/domain"
)

type ContainerApplication interface {
	FindOne(accountId string, containerId string) (*ContainerWithObjects, error)
	Save(container *Container) error
	Delete(accountId string, containerId string) error
	FindByAccount(accountId string) ([]*Container, error)
}

type ContainerWithObjects struct {
	Container *Container
	Objects   []*Object
}

func ContainerWithObjectsFrom(container *domain.Container, domainObjects []*domain.Object) *ContainerWithObjects {
	var objects []*Object
	for _, object := range domainObjects {
		objects = append(objects, ObjectFrom(object))
	}
	return &ContainerWithObjects{
		Container: ContainerFrom(container),
		Objects:   objects,
	}
}

type DefaultContainerApplication struct {
	accountRepository   domain.AccountRepository
	containerRepository domain.ContainerRepository
	objectRepository    domain.ObjectRepository
}

func NewDefaultContainerApplication(accountRepository domain.AccountRepository, containerRepository domain.ContainerRepository, objectRepository domain.ObjectRepository) *DefaultContainerApplication {
	return &DefaultContainerApplication{accountRepository: accountRepository, containerRepository: containerRepository, objectRepository: objectRepository}
}

func (d *DefaultContainerApplication) FindOne(accountId string, containerId string) (*ContainerWithObjects, error) {
	account, err := d.accountRepository.FindOne(accountId)
	if err != nil {
		return nil, domain.NewNotFoundError(fmt.Sprintf("not found container '%s/%s'", accountId, containerId), err)
	}
	container, err := d.containerRepository.FindOne(containerId, account)
	if err != nil {
		return nil, err
	}
	objects, err := d.objectRepository.FindByContainer(container)
	if err != nil {
		return nil, err
	}
	return ContainerWithObjectsFrom(container, objects), nil
}

func (d *DefaultContainerApplication) Save(container *Container) error {
	panic("implement me")
}

func (d *DefaultContainerApplication) Delete(accountId string, containerId string) error {
	panic("implement me")
}

func (d *DefaultContainerApplication) FindByAccount(accountId string) ([]*Container, error) {
	panic("implement me")
}
