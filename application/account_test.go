package application

import (
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountApplication_Find(t *testing.T) {
	accountRepository := &mockAccountRepository{}
	accountApp := NewAccountApplication(accountRepository)

	accountRepository.On("Find", "reuben").Return(domain.NewAccount("reuben"), nil)
	accountRepository.On("Find", "jimmy").Return(nil, domain.NewNotFoundError("cannot find", nil))

	t.Run("find account", func(t *testing.T) {
		reuben, err := accountApp.Find("reuben")
		expected := NewAccount("reuben")
		assert.Nil(t, err)
		assert.Equal(t, expected, reuben)
	})

	t.Run("cannot find account", func(t *testing.T) {
		jimmy, err := accountApp.Find("jimmy")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.IsType(t, expected, err)
		assert.Nil(t, jimmy)
	})

	accountRepository.AssertExpectations(t)
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
	panic("implement me")
}
