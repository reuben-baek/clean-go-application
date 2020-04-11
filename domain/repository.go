package domain

type AccountRepository interface {
	FindOne(id string) (*Account, error)
	Save(account *Account) error
	Delete(account *Account) error
}

type ContainerRepository interface {
	FindOne(id string, account *Account) (*Container, error)
	FindByAccount(account *Account) ([]*Container, error)
	Save(container *Container) error
	Delete(container *Container) error
}

type ObjectRepository interface {
	FindOne(id string, container *Container) (*Object, error)
	FindByContainer(container *Container) ([]*Object, error)
	Create(id string, container *Container) (*Object, error)
	Save(object *Object) error
	Delete(object *Object) error
}

type NotFoundError struct {
	message string
	err     error
}

func NewNotFoundError(message string, err error) *NotFoundError {
	return &NotFoundError{message: message, err: err}
}

func (n *NotFoundError) Error() string {
	if n.err != nil {
		return "NotFoundError: " + n.message + "; " + n.err.Error()
	} else {
		return "NotFoundError: " + n.message
	}
}

func (n *NotFoundError) Unwrap() error {
	return n.err
}
