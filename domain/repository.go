package domain

type AccountRepository interface {
	Find(id string) (*Account, error)
	Save(account *Account) error
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
