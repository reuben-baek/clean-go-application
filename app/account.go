package app

import "github.com/reuben-baek/clean-go-application/domain"

type Account struct {
	Id string
}

func (a *Account) To() *domain.Account {
	return domain.NewAccount(a.Id)
}

func AccountFrom(account *domain.Account) Account {
	return Account{account.Id()}
}

type Container struct {
	Id string
}

func ContainerFrom(container *domain.Container) Container {
	return Container{Id: container.Id()}
}

type AccountWithContainers struct {
	Account    Account
	Containers []Container
}

func AccountWithContainersFrom(account *domain.Account, domainContainers []*domain.Container) AccountWithContainers {
	var containers []Container
	for _, c := range domainContainers {
		containers = append(containers, ContainerFrom(c))
	}
	return AccountWithContainers{
		Account:    AccountFrom(account),
		Containers: containers,
	}
}

type AccountApplication interface {
	Find(id string) (AccountWithContainers, error)
	Save(account Account) error
	Delete(id string) error
}

type DefaultAccountApplication struct {
	accountRepository   domain.AccountRepository
	containerRepository domain.ContainerRepository
}

func NewDefaultAccountApplication(accountRepository domain.AccountRepository, containerRepository domain.ContainerRepository) *DefaultAccountApplication {
	return &DefaultAccountApplication{accountRepository: accountRepository, containerRepository: containerRepository}
}

func (d *DefaultAccountApplication) Find(id string) (AccountWithContainers, error) {
	account, err := d.accountRepository.FindOne(id)
	if err != nil {
		return AccountWithContainers{}, err
	}
	containers, err := d.containerRepository.FindByAccount(account)
	if err != nil {
		return AccountWithContainers{}, err
	}
	return AccountWithContainersFrom(account, containers), nil
}

func (d *DefaultAccountApplication) Save(account Account) error {
	panic("implement me")
}

func (d *DefaultAccountApplication) Delete(id string) error {
	panic("implement me")
}
