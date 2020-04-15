package application

import (
	"errors"
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountApplication_Find(t *testing.T) {
	accountRepository := &accountRepository{}
	containerRepository := &containerRepository{}

	reuben := domain.NewAccount("reuben")
	reubenContainers := []*domain.Container{
		domain.NewContainer("document", reuben),
	}
	accountRepository.On("FindOne", "reuben").Return(reuben, nil)
	accountRepository.On("FindOne", "jimmy").Return(nil, domain.NewNotFoundError("not found", nil))

	containerRepository.On("FindByAccount", reuben).Return(reubenContainers, nil)
	accountApp := NewDefaultAccountApplication(accountRepository, containerRepository)

	t.Run("found", func(t *testing.T) {
		reubenContainerList, err := accountApp.FindOne("reuben")
		expected := &AccountWithContainers{
			Account: NewAccount("reuben"),
			Containers: []*Container{
				NewContainer("document"),
			},
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, reubenContainerList)
	})

	t.Run("not found error", func(t *testing.T) {
		jimmy, err := accountApp.FindOne("jimmy")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.IsType(t, expected, err)
		assert.Nil(t, jimmy)
	})
}

func TestAccountApplication_Save(t *testing.T) {
	accountRepository := &accountRepository{}
	containerRepository := &containerRepository{}

	accountRepository.On("Save", domain.NewAccount("bob")).Return(nil)
	accountRepository.On("Save", domain.NewAccount("ted")).Return(errors.New("unexpected error"))
	accountApp := NewDefaultAccountApplication(accountRepository, containerRepository)

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

func TestAccountApplication_Delete(t *testing.T) {
	accountRepository := &accountRepository{}
	containerRepository := &containerRepository{}

	accountRepository.On("FindOne", "reuben").Return(domain.NewAccount("reuben"), nil)
	accountRepository.On("Delete", domain.NewAccount("reuben")).Return(nil)
	accountRepository.On("FindOne", "bob").Return(domain.NewAccount("bob"), nil)
	accountRepository.On("Delete", domain.NewAccount("bob")).Return(errors.New("unexpected error"))
	accountRepository.On("FindOne", "ted").Return(nil, domain.NewNotFoundError("not found", nil))
	accountApp := NewDefaultAccountApplication(accountRepository, containerRepository)

	t.Run("success", func(t *testing.T) {
		err := accountApp.Delete("reuben")
		assert.Nil(t, err)
	})
	t.Run("not found", func(t *testing.T) {
		err := accountApp.Delete("ted")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.IsType(t, expected, err)
	})
	t.Run("unexpected error", func(t *testing.T) {
		err := accountApp.Delete("bob")
		assert.NotNil(t, err)
	})
}

type accountRepository struct {
	mock.Mock
}

func (r *accountRepository) FindOne(id string) (*domain.Account, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Account), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (r *accountRepository) Save(account *domain.Account) error {
	args := r.Called(account)
	return args.Error(0)
}

func (r *accountRepository) Delete(account *domain.Account) error {
	args := r.Called(account)
	return args.Error(0)
}

type containerRepository struct {
	mock.Mock
}

func (r *containerRepository) FindOne(id string, account *domain.Account) (*domain.Container, error) {
	args := r.Called(id, account)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Container), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (r *containerRepository) FindByAccount(account *domain.Account) ([]*domain.Container, error) {
	args := r.Called(account)
	return args.Get(0).([]*domain.Container), args.Error(1)
}

func (r *containerRepository) Save(account *domain.Container) error {
	args := r.Called(account)
	return args.Error(0)
}

func (r *containerRepository) Delete(account *domain.Container) error {
	args := r.Called(account)
	return args.Error(0)
}
