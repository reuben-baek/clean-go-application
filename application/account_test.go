package application_test

import (
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountApplication_Find(t *testing.T) {
	accountRepository := &AccountRepository{}
	accountApp := application.NewAccountApplication(accountRepository)

	t.Run("find account", func(t *testing.T) {
		reuben, err := accountApp.Find("reuben")
		expected := application.NewAccount("reuben")
		assert.Nil(t, err)
		assert.Equal(t, expected, reuben)
	})

	t.Run("cannot find account", func(t *testing.T) {
		jimmy, err := accountApp.Find("jimmy'")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.IsType(t, expected, err)
		assert.Nil(t, jimmy)
	})
}

type AccountRepository struct {
	mock.Mock
}

func (a *AccountRepository) Find(id string) (*domain.Account, error) {
	panic("implement me")
}

func (a *AccountRepository) Save(account *domain.Account) error {
	panic("implement me")
}

func (r *AccountRepository) Delete(account *domain.Account) error {
	panic("implement me")
}
