package domain

type Account interface {
	Id() string
}
type account struct {
	id string
}

func NewAccount(id string) Account {
	return &account{id: id}
}

func (a *account) Id() string {
	return a.id
}
