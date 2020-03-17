# Domain Definition

account 도메인 모델을 정의한다. account 는 object storage 의 소유자이다. 일단 id 필드만 정의한다. 

[account.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/account.go)

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

NewAccount 는 Account 생성자 이다. go 에서는 생성자를 사용하지 않고 struct type composite literal ( `Account{id: id}` ) 을 이용할 수 도 있지만, 필드나 생성 로직이 추가되는 경우 컴파일 단계에서 코드를 수정하고 오류를 해결할 수 있도록 생성자를 사용하는 것이 강건한 코드를 유지하는데 도움이 되기 때문에 되도록 생성자를 만들겠다.

Id() 는 Account.id 의 Getter 이다. id 필드명이 소문자로 시작하므로 외부에 Setter 가 열려있지 않다. id 는 한번 생성되면 불변이므로 Setter 를 오픈하지 않는다. account_test.go 를 보자. package 를 domain_test 로 해서 외부에서 Account 를 사용하도록 했다. id 필드는 노출되어 있지 않으므로 사용하려면 컴파일 오류가 발생한다. 

[account_test.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/account_test.go)

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

이제 account 객체를 저장하는 AccountRepository 를 정의한다. AccountRepository 는 db, filesystem 등 다양한 형태로 구현될 수 있으므로 interface 로 정의한다. 구현체들은 infrastructure layer 에서 다른 패키지로 존재하게 된다.

[repository.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/repository.go)

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

Account 도메인 모델 정의는 일단 끝났다. 다음 단계는 application layer 또는 infrastructure layer 구현이다. 어느 layer 나 먼저 진행해도 되고 동시에 진행 ( 팀 멤버가 두명 이상이라면 ) 해도 된다. application layer 와 infrastructure layer 는 서로 의존성이 없기 때문이다. 둘다 domain layer 에만 의존성을 가지고 있다. infrastructure layer 를 먼저 구현하면 application layer 구현시점에 테스트 코드를 infrastructure layer 에 의존하도록 할 가능성이 높기 때문에 ( Mocking 이 귀찮거나 모르거나 ) 여기서는 application layer 구현에 대해서 먼저 진행하겠다. ( 사실, 실제 코드 구현은 infrastructure layer 를 먼저했다. 안좋은 습관이긴 한데, 의존성 없도록 application layer 구현할 테니 상관은 없다. )

