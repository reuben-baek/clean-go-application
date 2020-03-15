package domain

type Account struct {
	id string
}

func NewAccount(id string) *Account {
	return &Account{id: id}
}

func (a *Account) Id() string {
	return a.id
}
