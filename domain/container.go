package domain

type Container struct {
	id      string
	account *Account
}

func NewContainer(id string, account *Account) *Container {
	return &Container{id: id, account: account}
}

func (c *Container) Id() string {
	return c.id
}

func (c *Container) Account() *Account {
	return c.account
}
