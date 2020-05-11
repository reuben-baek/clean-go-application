 이번장에서는 어플리케이션 레이어를 구현한다. 어플리케이션 레이어는 도메인 레이어에서 정의한 모델들 - 엔터티, 레파지토리, 서비스등 - 을 조합해서 어플리케이션 요구사항을 구현한다. 어플리케이션 구현시 주의할점은 인프라스트럭처 레이어에 의존성을 가지지 않도록 하는 것이다. 즉, 인프라스트럭처 레이어에 구현된 구현체 타입을 사용하면 안된다. 인프라스트럭처 레이어에 구현된 코드가 하나도 없다고 생각하고 어플리케이션 레이어를 구현하는 것도 좋은 방법이다. 

여기서 구현하는 어플리케이션은 Simple Web Object Storage 이다. Web 인터페이스는 인터페이스 레이어에서 구현할 예정이다. 인터페이스 레이어는 사용자의 요청을 받고 어플리케이션 레이어에 맞도록 요청을 변환하여 어플리케이션을 호출한다. 어플리케이션은 요청을 처리하고 결과를 리턴한다. 어플리케이션의 요구사항은 변경 가능성이 도메인 영역 보다 높다. 어플리케이션의 변경이 도메인 레이어에 영향을 주지 않도록 하려면 어플리케이션 호출 인터페이스가 도메인 모델을 사용하지 않도록 해야 한다. 흔히 DTO - Data Trasnfer Object - 라는 표현으로 사용되는 데이터 타입을 정의하고, 어플리케이션 레이어와 인터페이스 레이어 간의 상호 작용에 DTO 를 이용해야 한다. 이렇게 함으로서 인터페이스 레이어가 온전히 어플리케이션 레이어에만 의존성을 가지도록 할 수 있다.

[application_requirement.md]() 에 정의된 요구사항에 따라 Account, Container, Object 를 생성, 삭제, 변경, 조회 하는 어플리케이션을 TDD 로 작성해 보자.

## AccountApplication 구현

아래 Account Web 인터페이스 스펙으로 부터 AccountApplication 인터페이스를 만들어보자.

```
PUT /:account body:account-meta
- 신규 account를 생성한다.
GET /:account
- account 메타정보와 account의 container 목록을 제공한다.
POST /:account body:account-meta
- account 메타정보를 수정한다.
DELETE /:account
- account를 삭제한다
```

먼저, DTO 를 정의하자. PUT, POST 요청으로 부터 Account type - account id 와 메타 정보 - 을 정의할 수 있다. Get 응답으로 부터 Container, AccountWithContainers type 을 정의한다.

```go
package application

type Account struct {
	Id string
}

type Container struct {
  Id string
}

type AccountWithContainers struct {
	Account Account
	Containers []Container
}
```

위에서 정의한 DTO type 을 가지고 AccountApplication 인터페이스를 아래와 같이 정의한다.

```go
package application
type AccountApplication interface {
	Find(id string) (AccountWithContainers, error)
	Save(account Account) error
	Delete(id string) error
}
```

Web 인터페이스 Get 요청은 Find(id string) 을 호출하여 (AccountWithContainers, error) 를 받는다. 메소드 호출 및 리턴 타입으로 DTO pointer 를 사용하지 않고 value 를 사용하는 것으 눈여겨 보자. go 메소드 호출 파라미터는 메소드 호출시 copy 된다. pointer 를 사용하면 pointer 값이 copy 되고, value 를 사용하면 value 가 copy 된다. 객체 T type 에 대한 pointer 값을 전달하면 호출된 메소드는 해당 객체를 직접 사용하는 것이고, 객체에 대한 value 를 전달하면 호출된 메소드는 해당 객체의 copy 된 객체를 사용하는 것이다. 따라서, T type 객체가 내부 상태를 변경하는 메소드를 노출하거나 필드가 Public 하다면 호출된 메소드는 이를 이용해서 객체 내부상태를 변경할 수 있고 호출자도 변화된 객체를 보게된다 . 하지만, 객체의 copy 가 전달된다면 호출된 메소드에서 객체 상태를 변경한다 하더라도 호출자가 가지고 있는 객체의 상태는 호출 이전과 다르지않다.

함수형 프로그래밍의 관점에서, 함수 호출은 호출자에게 어떤 변화도 없이 결과값만을 받기를 기대한다. 즉, side effect 가 없는 순수 함수를 사용함으로서 얻는 이점은 매우 크다. 코드를 이해하기 쉽게 하고, 의도치 않는 버그가 발생할 확률을 줄여준다. 이런 측면에서 여기에서는 가능한 불변 객체를 사용하려 한다. 불변 객체는 두가지 형태로 사용된다. 객체 reference 를 사용하여 주고 받으면서도 객체의 상태 변경 메소드를 노출하지 않는 방식 과 파라미터 전달시 copy 된 객체를 전달하는 방식을 이용한다. DTO 의 경우는 후자를 사용하겠다.	



이제 AccountApplication 을 TDD 로 구현해 보자.

```
type DefaultAccountApplication struct {
}

func NewDefaultAccountApplication() *DefaultAccountApplication {
	return &DefaultAccountApplication{}
}

func (d *DefaultAccountApplication) Find(id string) (AccountWithContainers, error) {
	panic("implement me")
}

func (d *DefaultAccountApplication) Save(account Account) error {
	panic("implement me")
}

func (d *DefaultAccountApplication) Delete(id string) error {
	panic("implement me")
}
```



```
func TestDefaultAccountApplication_Find(t *testing.T) {
	accountApp := NewDefaultAccountApplication()
	accountWithContainers, err := accountApp.Find("reuben")
	expected := AccountWithContainers{
		Account:    Account{"reuben"},
		Containers: []Container{Container{"document"}},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, accountWithContainers)
}
```

테스트 실행하면,

```
$go test  github.com/reuben-baek/clean-go-application/application
	FAIL: TestDefaultAccountApplication_Find (0.00s)
panic: implement me [recovered]
        panic: implement me
goroutine 20 [running]:
testing.tRunner.func1(0xc00010a100)
        /usr/local/go/src/testing/testing.go:874 +0x3a3
panic(0x12b9f00, 0x138d4d0)
        /usr/local/go/src/runtime/panic.go:679 +0x1b2
github.com/reuben-baek/clean-go-application/application.(*DefaultAccountApplication).Find(...)
        /Users/reuben/ReubenProjects/clean-go-application/application/account.go:54
github.com/reuben-baek/clean-go-application/app.TestDefaultAccountApplication_Find(0xc00010a100)
        /Users/reuben/ReubenProjects/clean-go-application/application/account_test.go:10 +0x39
```



DefaultAccountApplication.Find 를 구현하자.

```
func (d *DefaultAccountApplication) Find(id string) (AccountWithContainers, error) {
	return AccountWithContainers{
		Account:    Account{"reuben"},
		Containers: []Container{Container{"document"}},
	}, nil
}
```

테스트 실행

```
$go test  github.com/reuben-baek/clean-go-application/application
ok      github.com/reuben-baek/clean-go-application/application 0.345s
```

오류 테스트 케이스 만들자. 

```
func TestDefaultAccountApplication_Find(t *testing.T) {
	accountApp := NewDefaultAccountApplication()

	t.Run("found", func(t *testing.T) {
		accountWithContainers, err := accountApp.Find("reuben")
		expected := AccountWithContainers{
			Account:    Account{"reuben"},
			Containers: []Container{Container{"document"}},
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, accountWithContainers)
	})

	t.Run("not found", func(t *testing.T) {
		accountWithContainers, err := accountApp.Find("bob")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.Nil(t, accountWithContainers)
		assert.Equal(t, expected, err)
	})
}
```

테스트 실행

```
--- FAIL: TestDefaultAccountApplication_Find (0.00s)
    --- FAIL: TestDefaultAccountApplication_Find/not_found (0.00s)
        account_test.go:25: 
                Error Trace:    account_test.go:25
                Error:          Expected nil, but got: application.AccountWithContainers{Account:application.Account{Id:"reuben"}, Containers:[]application.Container{application.Container{Id:"document"}}}
                Test:           TestDefaultAccountApplication_Find/not_found
        account_test.go:26: 
                Error Trace:    account_test.go:26
                Error:          Not equal: 
                                expected: *domain.NotFoundError(&domain.NotFoundError{message:"cannot find", err:error(nil)})
                                actual  : <nil>(<nil>)
                Test:           TestDefaultAccountApplication_Find/not_found
FAIL
FAIL    github.com/reuben-baek/clean-go-application/application 0.345s
FAIL

```

not found 테스트케이스 를 통과 시키려면, 등록된 account 가 아니면 에러를 내야 한다. 등록된 account 여부를 확인하려면, domain.AccountRepository 를 호출해서 account 정보를 가져오면 된다.

```
type DefaultAccountApplication struct {
	accountRepository domain.AccountRepository
}

func NewDefaultAccountApplication(accountRepository domain.AccountRepository) *DefaultAccountApplication {
	return &DefaultAccountApplication{accountRepository: accountRepository}
}

func (d *DefaultAccountApplication) Find(id string) (AccountWithContainers, error) {
	account, err := d.accountRepository.FindOne(id)
	if err != nil {
		return AccountWithContainers{}, err
	}
	return AccountWithContainers{
		Account:    AccountFrom(account),
		Containers: []Container{Container{"document"}},
	}, nil
}
```

AccountFrom 과 같은 DTO helper function 들을 구현해 놓으면 코드가 좀더 깔끔해진다.

```
func AccountFrom(account *domain.Account) Account {
	return Account{account.Id()}
}
func ContainerFrom(container *domain.Container) Container {
	return Container{Id: container.Id()}
}
func AccountWithContainersFrom(account *domain.Account, domainContainers []*domain.Container) AccountWithContainers {
	var containers []Container
	for _, c := range domainContainers {
		containers = append(containers, ContainerFrom(c))
	}
	return AccountWithContainers{
		Account:    AccountFrom(account),
		Containers: containers,
	}
}
```

accountRepository 구현체가 필요하다. mock 으로 구현한다.

```
func TestDefaultAccountApplication_Find(t *testing.T) {
	accountRepository := new(accountRepository)
	accountApp := NewDefaultAccountApplication(accountRepository)

	t.Run("found", func(t *testing.T) {
		accountWithContainers, err := accountApp.Find("reuben")
		expected := AccountWithContainers{
			Account:    Account{"reuben"},
			Containers: []Container{Container{"document"}},
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, accountWithContainers)
	})

	t.Run("not found", func(t *testing.T) {
		accountRepository.On("FindOne", "bob").Return((*domain.Account)(nil), domain.NewNotFoundError("cannot find", nil))
		accountWithContainers, err := accountApp.Find("bob")
		expected := domain.NewNotFoundError("cannot find", nil)
		assert.Equal(t, AccountWithContainers{}, accountWithContainers)
		assert.Equal(t, expected, err)
	})
}

type accountRepository struct {
	mock.Mock
}

func (a *accountRepository) FindOne(id string) (*domain.Account, error) {
	args := a.Called(id)
	return args.Get(0).(*domain.Account), args.Error(1)
}
```

테스트를 실행시키면, not found 케이스는 성공한다. 하지만, found 실패한다. found 케이스에 성공시 mock 설정해준다.



```
	t.Run("found", func(t *testing.T) {
		accountRepository.On("FindOne", "reuben").Return(domain.NewAccount("reuben"), nil)
		accountWithContainers, err := accountApp.Find("reuben")
		expected := AccountWithContainers{
			Account:    Account{"reuben"},
			Containers: []Container{Container{"document"}},
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, accountWithContainers)
	})
```

성공한다. 새로운 테스트 케이스 하나 더 만들자.  container 를 가지고 있지 않은 jimmy 를 찾아보자.

```
	t.Run("found jimmy", func(t *testing.T) {
		accountRepository.On("FindOne", "jimmy").Return(domain.NewAccount("jimmy"), nil)
		accountWithContainers, err := accountApp.Find("jimmy")
		expected := AccountWithContainers{
			Account:    Account{"jimmy"},
			Containers: []Container(nil),
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, accountWithContainers)
	})
```

테스트 실패

```
Expected :app.AccountWithContainers{Account:app.Account{Id:"jimmy"}, Containers:[]app.Container{}}
Actual   :app.AccountWithContainers{Account:app.Account{Id:"jimmy"}, Containers:[]app.Container{app.Container{Id:"document"}}}
```

다시 DefaultAccountApplication.Find 로 돌아가서 account 로 container 목록을 가져오는 기능을 추가하자.

```
type DefaultAccountApplication struct {
	accountRepository domain.AccountRepository
	containerRepository domain.ContainerRepository
}

func NewDefaultAccountApplication(accountRepository domain.AccountRepository, containerRepository domain.ContainerRepository) *DefaultAccountApplication {
	return &DefaultAccountApplication{accountRepository: accountRepository, containerRepository: containerRepository}
}


func (d *DefaultAccountApplication) Find(id string) (AccountWithContainers, error) {
	account, err := d.accountRepository.FindOne(id)
	if err != nil {
		return AccountWithContainers{}, err
	}
	containers, err := d.containerRepository.FindByAccount(account)
	if err != nil {
		return AccountWithContainers{}, err
	}
	return AccountWithContainersFrom(account, containers), nil
}
```

테스트 수정해야 한다. containerRepository mock 을 추가하고 인스턴스를 만들어서 NewDefaultAccountApplication 에 주입한다.

```
func TestDefaultAccountApplication_Find(t *testing.T) {
	accountRepository := new(accountRepository)
	containerRepository := new(containerRepository)
	accountApp := NewDefaultAccountApplication(accountRepository, containerRepository)
	...
}

type containerRepository struct {
	mock.Mock
}

func (c *containerRepository) FindOne(id string, account *domain.Account) (*domain.Container, error) {
}

func (c *containerRepository) FindByAccount(account *domain.Account) ([]*domain.Container, error) {
	args := c.Called(account)
	return args.Get(0).([]*domain.Container), args.Error(1)}

func (c *containerRepository) Save(container *domain.Container) error {
	panic("implement me")
}

func (c *containerRepository) Delete(container *domain.Container) error {
	panic("implement me")
}
```

테스트 실행하면, containerRepository.FindByAccount 에 대한 mock 정의가 필요하다는 오류가 나온다.

```
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
```

테스트 실행하면, 성공한다. found reuben 케이스도 containerRepository 에 대한 mock 설정을 해 주면 된다.

```
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
```

모든 테스트 케이스가 성공한다.

```
$ go test  github.com/reuben-baek/clean-go-application/application
ok      github.com/reuben-baek/clean-go-application/application 0.693s

```

DefaultAccountApplication.Find 구현이 완성되었다.

Save, Delete 구현도 해 보자.



---

## ContainerApplication 구현

