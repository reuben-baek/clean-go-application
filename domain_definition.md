# Domain Definition

Simple Object Storage 어플리케이션 요구사항에서 도메인 모델을 정의 해 보자. 일단 account, container, object 가 필요해 보인다.

## Account

account 는 object storage 의 소유자이다. 일단, 속성으로 id 필드만 정의한다. 

[domain/account.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/account.go)

```go
type Account struct {
	id string
}

func NewAccount(id string) *Account {
	return &Account{id: id}
}

func (a *Account) Id() string {
	return a.id
}
```

Account 의 속성 명을 외부 패키지에서 참조 할 수 없도록 소문자 - id 로 만들었다. Getter function 으로 Id() 를 제공한다. id 값은 한 번 만들어지면 변하지 않으므로 Setter 는 필요없다. 전통적인 OOP 에서의 클래스 캡슐화를 go 에서 구현하는 방법이다. 

Account 생성자로 NewAccount(id string) 을 제공한다. 외부 패키지에서 NewAccount() 뿐 아니라 struct type composite literal - `&Account{"reuben"}` - 을 사용해서 Account 객체를 생성할 수 도 있다. 그럼, Account type 도 굳이 외부 패키지에 오픈할 이유는 없지 않을까? composite literal 로 객체를 만들 수 없도록 강제하려면 type 명을 `type account struct{}` 소문자로 시작하게 하면 된다. 그러나, type 은 composite literal 에서만 사용되는게 아니라 `var reuben Account` 와 같이 type 선언 및 type casting ` v.(*Account)` 에도 필요하다. 외부 패키지에서 이와 같이 사용하려면 type 명을 대문자로 시작하게 해서 외부 패키지에 오픈해야 한다. 그래서, 실용적인 선택으로 여기서는 타입명은 대문자로 시작해서 외부에 오픈하고, 속성은 소문자로 시작해서 외부에 오픈하지 않고 Getter 를 오픈하도록 하겠다. 필요한 경우 SetId(id string) 과 같이 Setter 를 정의하겠지만, 꼭 필요한 경우에 한정해서 사용하겠다.

account_test.go 를 보자. package 를 domain_test 로 해서 외부에서 Account 를 사용하도록 했다. id 필드는 노출되어 있지 않으므로 사용하려면 컴파일 오류가 발생한다. 

[domain/account_test.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/account_test.go)

```go
package domain_test

import (
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {
  var reuben *domain.Account
	reuben = domain.NewAccount("reuben")
	assert.Equal(t, "reuben", reuben.Id())
  // p := reuben.id //  compile error : Unexported field 'id' usage
}
```

assert library 로 go community 에서 가장 많이 사용하는 github.com/stretchr/testify 를 추가했다. 외부 모듈 관리 방법으로 이 프로젝트에서는 go module 을 사용한다.

```
go mod init github.com/reuben-baek/clean-go-application
go get github.com/stretchr/testify
```

struct type composite liternal 대신 생성자를 사용함으로서 얻는 큰 이익이 있다. 속성, 생성 로직에 변경이 발생하는 경우 컴파일 단계에서 코드를 수정하고 오류를 해결할 수 있다. Account type 에 name 속성을 추가하고 NewAccount 생성자를 변경해보자.

```go
type Account struct {
	id string
	name string
}

func NewAccount(id string, name string) *Account {
	return &Account{id: id}
}
```

NewAccount 를 사용한 기존 코드는 컴파일 오류가 발생한다. 코드를 실행하려면, 변경된 NewAccount 에 맞게 모두 코드를 수정해야 한다.

```go
reuben = domain.NewAccount("reuben") // not enough arguments in call to domain.NewAccount
```

반면, type compisite literal 을 이용해 객체를 생성한 아래 코드는 컴파일 오류가 발생하지 않는다. 지금 단계에서는 별 의미가 없을 수도 있지만 runtime 에 문제가 발생할 가능성이 매우 높아진다.

```go
func TestAccount_NewWithLiteral(t *testing.T) {
	reuben := &domain.Account{"reuben"}
	assert.Equal(t, "", reuben.Id())
}
```



## Container





## Object



## AccountRepository

이제 account 객체를 저장하는 AccountRepository 를 정의한다. AccountRepository 는 db, filesystem 등 다양한 형태로 구현될 수 있으므로 interface 로 정의한다. 구현체들은 infrastructure layer 에서 다른 패키지로 존재하게 된다.

[domain/repository.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/repository.go)

```go
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
```

Find 메소드는 id 로 Account 를 찾지 못하는 경우에 NotFoundError 를 리턴한다. 구현체에 따라 NotFoundError 유발 오류가 다를 수 있어서 NewNotFoundError 생성자로 error 를 받도록 했다. AccountRepository 와 관련있는 에러이기 때문에 같은 소스(repository.go)에 정의한다. 



## ContainerRepository



## ObjectRepository



도메인 모델 정의는 일단 끝났다. 다음 단계는 application layer 또는 infrastructure layer 구현이다. 어느 layer 나 먼저 진행해도 되고 동시에 진행 ( 팀 멤버가 두명 이상이라면 ) 해도 된다. application layer 와 infrastructure layer 는 서로 의존성이 없기 때문이다. 둘다 domain layer 에만 의존성을 가지고 있다. infrastructure layer 를 먼저 구현하면 application layer 구현시점에 테스트 코드를 infrastructure layer 에 의존하도록 할 가능성이 높기 때문에 ( Mocking 이 귀찮거나 모르거나 ) 여기서는 application layer 구현에 대해서 먼저 진행하겠다. ( 사실, 실제 코드 구현은 infrastructure layer 를 먼저했다. 안좋은 습관이긴 한데, 의존성 없도록 application layer 구현할 테니 상관은 없다. )

