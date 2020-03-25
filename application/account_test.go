package application

import (
	"errors"
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountApplication_Find(t *testing.T) {
	accountRepository := &mockAccountRepository{}

	accountRepository.On("Find", "reuben").Return(domain.NewAccount("reuben"), nil)
	accountRepository.On("Find", "jimmy").Return(nil, domain.NewNotFoundError("cannot find", nil))

	accountApp := NewDefaultAccountApplication(accountRepository)

	t.Run("found", func(t *testing.T) {
		reuben, err := accountApp.Find("reuben")
		expected := NewAccount("reuben")
		assert.Nil(t, err)
		assert.Equal(t, expected, reuben)
	})

	t.Run("not found error", func(t *testing.T) {
		jimmy, err := accountApp.Find("jimmy")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.IsType(t, expected, err)
		assert.Nil(t, jimmy)
	})
}

func TestAccountApplication_Save(t *testing.T) {
	accountRepository := &mockAccountRepository{}

	accountRepository.On("Save", domain.NewAccount("bob")).Return(nil)
	accountRepository.On("Save", domain.NewAccount("ted")).Return(errors.New("unexpected error"))
	accountApp := NewDefaultAccountApplication(accountRepository)

	t.Run("success", func(t *testing.T) {
		bob := NewAccount("bob")
		err := accountApp.Save(bob)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ted := NewAccount("ted")
		err := accountApp.Save(ted)
		assert.NotNil(t, err)
	})
}

type mockAccountRepository struct {
	mock.Mock
}

func (r *mockAccountRepository) Find(id string) (*domain.Account, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Account), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (r *mockAccountRepository) Save(account *domain.Account) error {
	args := r.Called(account)
	return args.Error(0)
}
