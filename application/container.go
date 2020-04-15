package application

import (
	"fmt"
	"github.com/reuben-baek/clean-go-application/domain"
)

type ContainerApplication interface {
	FindOne(accountId string, containerId string) (*Container, error)
	Save(container *Container) error
	Delete(accountId string, containerId string) error
	FindByAccount(accountId string) ([]*Container, error)
}

type Container struct {
	Id string
}

func NewContainer(id string) *Container {
	return &Container{id}
}

func ContainerFrom(container *domain.Container) *Container {
	return &Container{Id: container.Id()}
}

type DefaultContainerApplication struct {
	accountRepository   domain.AccountRepository
	containerRepository domain.ContainerRepository
}

func NewDefaultContainerApplication(accountRepository domain.AccountRepository, containerRepository domain.ContainerRepository) *DefaultContainerApplication {
	return &DefaultContainerApplication{accountRepository: accountRepository, containerRepository: containerRepository}
}

func (d *DefaultContainerApplication) FindOne(accountId string, containerId string) (*Container, error) {
	account, err := d.accountRepository.FindOne(accountId)
	if err != nil {
		return nil, domain.NewNotFoundError(fmt.Sprintf("not found container '%s/%s'", accountId, containerId), err)
	}
	container, err := d.containerRepository.FindOne(containerId, account)
	if err != nil {
		return nil, err
	}
	return ContainerFrom(container), nil
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
