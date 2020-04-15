package application

import (
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDefaultContainerApplication_FindOne(t *testing.T) {
	accountRepository := &accountRepository{}
	reuben := domain.NewAccount("reuben")
	accountRepository.On("FindOne", "reuben").Return(reuben, nil)
	notFoundErrorBob := domain.NewNotFoundError("not found account 'bob'", nil)
	accountRepository.On("FindOne", "bob").Return(nil, notFoundErrorBob)

	containerRepository := &containerRepository{}
	containerRepository.On("FindOne", "document", reuben).Return(domain.NewContainer("document", reuben), nil)
	notFoundErrorMusic := domain.NewNotFoundError("not found container 'reuben/music'", nil)
	containerRepository.On("FindOne", "music", reuben).Return(nil, notFoundErrorMusic)
	containerApp := NewDefaultContainerApplication(accountRepository, containerRepository)

	t.Run("success", func(t *testing.T) {
		expected := NewContainer("document")
		container, err := containerApp.FindOne("reuben", "document")
		assert.Nil(t, err)
		assert.Equal(t, expected, container)
	})

	t.Run("not found - account", func(t *testing.T) {
		container, err := containerApp.FindOne("bob", "document")
		expected := domain.NewNotFoundError("not found container 'bob/document'", domain.NewNotFoundError("not found account 'bob'", nil))
		assert.Nil(t, container)
		assert.Equal(t, expected, err)
	})

	t.Run("not found - container", func(t *testing.T) {
		container, err := containerApp.FindOne("reuben", "music")
		expected := domain.NewNotFoundError("not found container 'reuben/music'", nil)
		assert.Nil(t, container)
		assert.Equal(t, expected, err)
	})
}

func TestDefaultContainerApplication_FindByAccount(t *testing.T) {
}

func TestDefaultContainerApplication_Save(t *testing.T) {
}

func TestDefaultContainerApplication_Delete(t *testing.T) {
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
	panic("implement me")
}

func (r *containerRepository) Save(account *domain.Container) error {
	args := r.Called(account)
	return args.Error(0)
}

func (r *containerRepository) Delete(account *domain.Container) error {
	args := r.Called(account)
	return args.Error(0)
}
