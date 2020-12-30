package domain

type Container interface {
	Id() string
	Account() Account
}

type container struct {
	id      string
	account Account
}

func NewContainer(id string, account Account) Container {
	return &container{id: id, account: account}
}

func (c *container) Id() string {
	return c.id
}

func (c *container) Account() Account {
	return c.account
}
