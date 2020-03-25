package application

import "github.com/reuben-baek/clean-go-application/domain"

type AccountApplication interface {
	Find(id string) (*Account, error)
	Save(account *Account) error
}

type DefaultAccountApplication struct {
	accountRepository domain.AccountRepository
}

func NewDefaultAccountApplication(accountRepository domain.AccountRepository) *DefaultAccountApplication {
	return &DefaultAccountApplication{accountRepository: accountRepository}
}

func (app *DefaultAccountApplication) Find(id string) (*Account, error) {
	account, err := app.accountRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return AccountFrom(account), nil
}

func (app *DefaultAccountApplication) Save(account *Account) error {
	return app.accountRepository.Save(account.To())
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
