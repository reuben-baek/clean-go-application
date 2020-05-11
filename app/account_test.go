package app

import (
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDefaultAccountApplication_Find(t *testing.T) {
	accountRepository := new(accountRepository)
	containerRepository := new(containerRepository)
	accountApp := NewDefaultAccountApplication(accountRepository, containerRepository)

	t.Run("found reuben", func(t *testing.T) {
		reuben := domain.NewAccount("reuben")
		accountRepository.On("FindOne", "reuben").Return(reuben, nil)
		containerRepository.On("FindByAccount", reuben).Return([]*domain.Container{domain.NewContainer("document", reuben)}, nil)
		accountWithContainers, err := accountApp.Find("reuben")
		expected := AccountWithContainers{
			Account:    Account{"reuben"},
			Containers: []Container{Container{"document"}},
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, accountWithContainers)
	})

	t.Run("not found bob", func(t *testing.T) {
		accountRepository.On("FindOne", "bob").Return((*domain.Account)(nil), domain.NewNotFoundError("cannot find", nil))
		accountWithContainers, err := accountApp.Find("bob")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.Equal(t, AccountWithContainers{}, accountWithContainers)
		assert.Equal(t, expected, err)
	})

	t.Run("found jimmy", func(t *testing.T) {
		jimmy := domain.NewAccount("jimmy")
		accountRepository.On("FindOne", "jimmy").Return(jimmy, nil)
		containerRepository.On("FindByAccount", jimmy).Return([]*domain.Container{}, nil)
		accountWithContainers, err := accountApp.Find("jimmy")
		expected := AccountWithContainers{
			Account:    Account{"jimmy"},
			Containers: []Container(nil),
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, accountWithContainers)
	})
}

type accountRepository struct {
	mock.Mock
}

func (a *accountRepository) FindOne(id string) (*domain.Account, error) {
	args := a.Called(id)
	return args.Get(0).(*domain.Account), args.Error(1)
}

func (a *accountRepository) Save(account *domain.Account) error {
	panic("implement me")
}

func (a *accountRepository) Delete(account *domain.Account) error {
	panic("implement me")
}

type containerRepository struct {
	mock.Mock
}

func (c *containerRepository) FindOne(id string, account *domain.Account) (*domain.Container, error) {
	args := c.Called(id)
	return args.Get(0).(*domain.Container), args.Error(1)
}

func (c *containerRepository) FindByAccount(account *domain.Account) ([]*domain.Container, error) {
	args := c.Called(account)
	return args.Get(0).([]*domain.Container), args.Error(1)
}

func (c *containerRepository) Save(container *domain.Container) error {
	panic("implement me")
}

func (c *containerRepository) Delete(container *domain.Container) error {
	panic("implement me")
}
