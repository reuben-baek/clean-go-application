ContainerApplication 정의



container.go

```go
package application

type ContainerApplication interface {
	FindOne(accountId string, containerId string) (*Container, error)
	Save(container *Container) error
	Delete(accountId string, containerId string) error
	FindByAccount(accountId string) ([]*Container, error)
}

type Container struct {
	Id string
}

func NewContainer(id string) *Container {
	return &Container{id}
}

type DefaultContainerApplication struct {

}

func NewDefaultContainerApplication() *DefaultContainerApplication {
	return &DefaultContainerApplication{}
}

func (d *DefaultContainerApplication) FindOne(accountId string, containerId string) (*Container, error) {
	panic("implement me")
}

func (d *DefaultContainerApplication) Save(container *Container) error {
	panic("implement me")
}

func (d *DefaultContainerApplication) Delete(accountId string, containerId string) error {
	panic("implement me")
}

func (d *DefaultContainerApplication) FindByAccount(accountId string) ([]*Container, error) {
	panic("implement me")
}


```



​	container_test.go

```go
package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultContainerApplication_FindOne(t *testing.T) {
	containerApp := NewDefaultContainerApplication()

	expected := NewContainer("document")
	container, err := containerApp.FindOne("reuben", "document")
	assert.Nil(t, err)
	assert.Equal(t, expected, container)
}
```



test fail

```
$go test github.com/reuben-baek/clean-go-application/application
--- FAIL: TestDefaultContainerApplication_FindOne (0.00s)
panic: implement me [recovered]
        panic: implement me

goroutine 7 [running]:
testing.tRunner.func1(0xc0000ee100)
        /usr/local/go/src/testing/testing.go:874 +0x3a3
panic(0x1379d20, 0x14840a0)
        /usr/local/go/src/runtime/panic.go:679 +0x1b2
github.com/reuben-baek/clean-go-application/application.(*DefaultContainerApplication).FindOne(...)
        /Users/reuben/ReubenProjects/clean-go-application/application/container.go:27
github.com/reuben-baek/clean-go-application/application.TestDefaultContainerApplication_FindOne(0xc0000ee100)
        /Users/reuben/ReubenProjects/clean-go-application/application/container_test.go:12 +0x61
testing.tRunner(0xc0000ee100, 0x1427e50)
        /usr/local/go/src/testing/testing.go:909 +0xc9
created by testing.(*T).Run
        /usr/local/go/src/testing/testing.go:960 +0x350
FAIL    github.com/reuben-baek/clean-go-application/application 0.699s
FAIL

```



```go
func (d *DefaultContainerApplication) FindOne(accountId string, containerId string) (*Container, error) {
	return NewContainer(containerId), nil
}
```



```
$go test github.com/reuben-baek/clean-go-application/application
ok      github.com/reuben-baek/clean-go-application/application 0.346s

```



```go
func TestDefaultContainerApplication_FindOne(t *testing.T) {
	containerApp := NewDefaultContainerApplication()

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
```



```
$go test github.com/reuben-baek/clean-go-application/application
--- FAIL: TestDefaultContainerApplication_FindOne (0.00s)
    --- FAIL: TestDefaultContainerApplication_FindOne/not_found_-_account (0.00s)
        container_test.go:22: 
                Error Trace:    container_test.go:22
                Error:          Expected nil, but got: &application.Container{Id:"document"}
                Test:           TestDefaultContainerApplication_FindOne/not_found_-_account
        container_test.go:23: 
                Error Trace:    container_test.go:23
                Error:          Not equal: 
                                expected: *domain.NotFoundError(&domain.NotFoundError{message:"not found container 'bob/document'", err:(*domain.NotFoundError)(0xc0000ae5a0)})
                                actual  : <nil>(<nil>)
                Test:           TestDefaultContainerApplication_FindOne/not_found_-_account
    --- FAIL: TestDefaultContainerApplication_FindOne/not_found_-_container (0.00s)
        container_test.go:29: 
                Error Trace:    container_test.go:29
                Error:          Expected nil, but got: &application.Container{Id:"music"}
                Test:           TestDefaultContainerApplication_FindOne/not_found_-_container
        container_test.go:30: 
                Error Trace:    container_test.go:30
                Error:          Not equal: 
                                expected: *domain.NotFoundError(&domain.NotFoundError{message:"not found container 'reuben/music'", err:(*domain.NotFoundError)(0xc0000ae720)})
                                actual  : <nil>(<nil>)
                Test:           TestDefaultContainerApplication_FindOne/not_found_-_container
FAIL
FAIL    github.com/reuben-baek/clean-go-application/application 0.334s
FAIL

```



```
type DefaultContainerApplication struct {
	accountRepository domain.AccountRepository
}

func NewDefaultContainerApplication(accountRepository domain.AccountRepository) *DefaultContainerApplication {
	return &DefaultContainerApplication{accountRepository}
}

func (d *DefaultContainerApplication) FindOne(accountId string, containerId string) (*Container, error) {
	_, err := d.accountRepository.FindOne(accountId)
	if err != nil {
		return nil, domain.NewNotFoundError(fmt.Sprintf("not found container '%s/%s'", accountId, containerId), err)
	}
	return NewContainer(containerId), nil
}
```



```
func TestDefaultContainerApplication_FindOne(t *testing.T) {
	accountRepository := &accountRepository{}
	containerApp := NewDefaultContainerApplication(accountRepository)
...
}
```



test "success"

```
panic: 
assert: mock: I don't know what to return because the method call was unexpected.
	Either do Mock.On("FindOne").Return(...) first, or remove the FindOne() call.
	This method was unexpected:
		FindOne(string)
		0: "reuben"
	at: [account_test.go:84 container.go:32 container_test.go:18] [recovered]
	panic: 
assert: mock: I don't know what to return because the method call was unexpected.
	Either do Mock.On("FindOne").Return(...) first, or remove the FindOne() call.
	This method was unexpected:
		FindOne(string)
		0: "reuben"
	at: [account_test.go:84 container.go:32 container_test.go:18]
```



```
func TestDefaultContainerApplication_FindOne(t *testing.T) {
  accountRepository := &accountRepository{}
	accountRepository.On("FindOne", "reuben").Return(domain.NewAccount("reuben"), nil)
	containerApp := NewDefaultContainerApplication(accountRepository)
...	
}
```



test "success" => pass

test "not found - account"

```
panic: 

mock: Unexpected Method Call
-----------------------------

FindOne(string)
		0: "bob"

The closest call I have is: 

FindOne(string)
		0: "reuben"


Diff: 0: FAIL:  (string=bob) != (string=reuben) [recovered]
```



```
func TestDefaultContainerApplication_FindOne(t *testing.T) {
  accountRepository := &accountRepository{}
	accountRepository.On("FindOne", "reuben").Return(domain.NewAccount("reuben"), nil)
	notFoundErrorBob := domain.NewNotFoundError("not found account 'bob'", nil)
	accountRepository.On("FindOne", "bob").Return(nil, notFoundErrorBob)
	containerApp := NewDefaultContainerApplication(accountRepository)
...	
}
```

test "not found - account" => pass



test "not found - container"

```
container_test.go:33: 
            	Error Trace:	container_test.go:33
            	Error:      	Expected nil, but got: &application.Container{Id:"music"}
            	Test:       	TestDefaultContainerApplication_FindOne/not_found_-_container
container_test.go:34: 
            	Error Trace:	container_test.go:34
            	Error:      	Not equal: 
            	            	expected: *domain.NotFoundError(&domain.NotFoundError{message:"not found container 'reuben/music'", err:error(nil)})
            	            	actual  : <nil>(<nil>)
            	Test:       	TestDefaultContainerApplication_FindOne/not_found_-_container


Expected :*domain.NotFoundError(&domain.NotFoundError{message:"not found container 'reuben/music'", err:error(nil)})
Actual   :<nil>(<nil>)
```



```
func ContainerFrom(container *domain.Container) *Container {
	return &Container{Id: container.Id()}
}

type DefaultContainerApplication struct {
	accountRepository domain.AccountRepository
	containerRepository domain.ContainerRepository
}

func NewDefaultContainerApplication(accountRepository domain.AccountRepository, containerRepository domain.ContainerRepository) *DefaultContainerApplication {
	return &DefaultContainerApplication{accountRepository: accountRepository, containerRepository: containerRepository}
}

func (d *DefaultContainerApplication) FindOne(accountId string, containerId string) (*Container, error) {
	account, err := d.accountRepository.FindOne(accountId)
	if err != nil {
		return nil, domain.NewNotFoundError(fmt.Sprintf("not found container '%s/%s'", accountId, containerId), err)
	}
	container, err := d.containerRepository.FindOne(containerId, account)
	if err != nil {
		return nil, err
	}
	return ContainerFrom(container), nil
}
```



```
$go test github.com/reuben-baek/clean-go-application/application
# github.com/reuben-baek/clean-go-application/application [github.com/reuben-baek/clean-go-application/application.test]
application/container_test.go:14:48: not enough arguments in call to NewDefaultContainerApplication
        have (*accountRepository)
        want (domain.AccountRepository, domain.ContainerRepository)
FAIL    github.com/reuben-baek/clean-go-application/application [build failed]

```



add mock containerRepository

```
type containerRepository struct {
	mock.Mock
}

func (r *containerRepository) FindOne(id string, account *domain.Account) (*domain.Container, error) {
	args := r.Called(id, account)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Container), args.Error(1)
	} else {
		return nil, args.Error(1)
	}}

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
```



fix error

```
func TestDefaultContainerApplication_FindOne(t *testing.T) {
	accountRepository := &accountRepository{}
	accountRepository.On("FindOne", "reuben").Return(domain.NewAccount("reuben"), nil)
	notFoundErrorBob := domain.NewNotFoundError("not found account 'bob'", nil)
	accountRepository.On("FindOne", "bob").Return(nil, notFoundErrorBob)

	containerRepository := &containerRepository{}
	containerApp := NewDefaultContainerApplication(accountRepository, containerRepository)
```



test "success"

```
panic: 
assert: mock: I don't know what to return because the method call was unexpected.
	Either do Mock.On("FindOne").Return(...) first, or remove the FindOne() call.
	This method was unexpected:
		FindOne(string,*domain.Account)
		0: "document"
		1: &domain.Account{id:"reuben"}
	at: [container_test.go:55 container.go:41 container_test.go:21] [recovered]
```



add mock on

```
func TestDefaultContainerApplication_FindOne(t *testing.T) {
	accountRepository := &accountRepository{}
	reuben := domain.NewAccount("reuben")
	accountRepository.On("FindOne", "reuben").Return(reuben, nil)
	notFoundErrorBob := domain.NewNotFoundError("not found account 'bob'", nil)
	accountRepository.On("FindOne", "bob").Return(nil, notFoundErrorBob)

	containerRepository := &containerRepository{}
	containerRepository.On("FindOne", "document", reuben).Return(domain.NewContainer("document", reuben), nil)
	containerApp := NewDefaultContainerApplication(accountRepository, containerRepository)
```

test "success" => pass

test "not found - account" => pass

test "not found - container"

```
mock: Unexpected Method Call
-----------------------------

FindOne(string,*domain.Account)
		0: "music"
		1: &domain.Account{id:"reuben"}

```

add mock on

```
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
```

test all

```
$ go test github.com/reuben-baek/clean-go-application/application
ok      github.com/reuben-baek/clean-go-application/application 0.346s
```

