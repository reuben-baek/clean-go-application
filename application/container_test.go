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
	reubenDocument := domain.NewContainer("document", reuben)
	containerRepository.On("FindOne", "document", reuben).Return(reubenDocument, nil)
	notFoundErrorMusic := domain.NewNotFoundError("not found container 'reuben/music'", nil)
	containerRepository.On("FindOne", "music", reuben).Return(nil, notFoundErrorMusic)

	objectRepository := &objectRepository{}
	reubenHello := domain.OpenObjectForRead("hello.txt", reubenDocument, 0, nil)
	objectRepository.On("FindByContainer", reubenDocument).Return([]*domain.Object{reubenHello}, nil)
	containerApp := NewDefaultContainerApplication(accountRepository, containerRepository, objectRepository)

	t.Run("success", func(t *testing.T) {
		expected := &ContainerWithObjects{
			Container: NewContainer("document"),
			Objects: []*Object{
				&Object{Id: "hello.txt"},
			},
		}
		containerWithObjects, err := containerApp.FindOne("reuben", "document")
		assert.Nil(t, err)
		assert.Equal(t, expected, containerWithObjects)
	})

	t.Run("not found - account", func(t *testing.T) {
		containerWithObjects, err := containerApp.FindOne("bob", "document")
		expected := domain.NewNotFoundError("not found container 'bob/document'", domain.NewNotFoundError("not found account 'bob'", nil))
		assert.Nil(t, containerWithObjects)
		assert.Equal(t, expected, err)
	})

	t.Run("not found - container", func(t *testing.T) {
		containerWithObjects, err := containerApp.FindOne("reuben", "music")
		expected := domain.NewNotFoundError("not found container 'reuben/music'", nil)
		assert.Nil(t, containerWithObjects)
		assert.Equal(t, expected, err)
	})
}

func TestDefaultContainerApplication_FindByAccount(t *testing.T) {
}

func TestDefaultContainerApplication_Save(t *testing.T) {
}

func TestDefaultContainerApplication_Delete(t *testing.T) {
}

type objectRepository struct {
	mock.Mock
}

func (o *objectRepository) FindOne(id string, container *domain.Container) (*domain.Object, error) {
	panic("implement me")
}

func (o *objectRepository) FindByContainer(container *domain.Container) ([]*domain.Object, error) {
	args := o.Called(container)
	return args.Get(0).([]*domain.Object), args.Error(1)
}

func (o *objectRepository) Create(id string, container *domain.Container) (*domain.Object, error) {
	panic("implement me")
}

func (o *objectRepository) Save(object *domain.Object) error {
	panic("implement me")
}

func (o *objectRepository) Delete(object *domain.Object) error {
	panic("implement me")
}
