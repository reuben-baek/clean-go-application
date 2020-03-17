package application

import "github.com/reuben-baek/clean-go-application/domain"

type AccountApplication struct {
	accountRepository domain.AccountRepository
}

func (app *AccountApplication) Find(id string) (*Account, error) {
	account, err := app.accountRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return AccountFrom(account), nil
}

func NewAccountApplication(accountRepository domain.AccountRepository) *AccountApplication {
	return &AccountApplication{accountRepository: accountRepository}
}

type Account struct {
	id string
}

func NewAccount(id string) *Account {
	return &Account{id: id}
}

func AccountFrom(account *domain.Account) *Account {
	return NewAccount(account.Id())
}
