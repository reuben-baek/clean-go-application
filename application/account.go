package application

import "github.com/reuben-baek/clean-go-application/domain"

type AccountApplication interface {
	FindOne(id string) (*AccountWithContainers, error)
	Save(account *Account) error
	Delete(id string) error
}

type DefaultAccountApplication struct {
	accountRepository   domain.AccountRepository
	containerRepository domain.ContainerRepository
}

func NewDefaultAccountApplication(accountRepository domain.AccountRepository, containerRepository domain.ContainerRepository) *DefaultAccountApplication {
	return &DefaultAccountApplication{accountRepository: accountRepository, containerRepository: containerRepository}
}

func (app *DefaultAccountApplication) FindOne(id string) (*AccountWithContainers, error) {
	account, err := app.accountRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	containers, err := app.containerRepository.FindByAccount(account)
	if err != nil {
		return nil, err
	}
	return AccountWithContainersFrom(account, containers), nil
}

func (app *DefaultAccountApplication) Save(account *Account) error {
	return app.accountRepository.Save(account.To())
}

func (app *DefaultAccountApplication) Delete(id string) error {
	account, err := app.accountRepository.FindOne(id)
	if err != nil {
		return err
	}
	return app.accountRepository.Delete(account)
}

type Account struct {
	Id string
}

func (a *Account) To() *domain.Account {
	return domain.NewAccount(a.Id)
}

func NewAccount(id string) *Account {
	return &Account{Id: id}
}

func AccountFrom(account *domain.Account) *Account {
	return NewAccount(account.Id())
}

type AccountWithContainers struct {
	Account    *Account
	Containers []*Container
}

func AccountWithContainersFrom(account *domain.Account, domainContainers []*domain.Container) *AccountWithContainers {
	var containers []*Container
	for _, c := range domainContainers {
		containers = append(containers, ContainerFrom(c))
	}
	return &AccountWithContainers{
		Account:    AccountFrom(account),
		Containers: containers,
	}
}
