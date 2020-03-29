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

account 는 여러 container 를 가질 수 있다. 즉, account 과 container 는 1:n 관계를 가진다. container 는 id 와 container 가 속한 account 를 속성으로 가지도록 정의한다. Getter 로 Id(), Account() function 을 노출한다.

[domain/container.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/container.go)

```go
type Container struct {
	id      string
	account *Account
}

func NewContainer(id string, account *Account) *Container {
	return &Container{id: id, account: account}
}

func (c *Container) Id() string {
	return c.id
}

func (c *Container) Account() *Account {
	return c.account
}
```

## Object

container 와 object 는 1:n 관계로 표현된다. account, container 관계와 같이 object 가 자신이 속한 container 를 속성으로 가지도록 정의한다.

```go
type Object struct {
	id        string
	container *Container
}
```

object 는 account, container 와 다르게 content 를 가지고 있다. 사용자는 임의의 데이터를 object.id 이름으로 업로드 하거나 다운받을 수 있다. object 가 content 를 가지고 있는 걸 어떻게 모델링 할까? 

먼저 생각해 볼 수 있는 방법은 `content []byte` 을 가지게 하는 것이다. object 의 content 를 object type 의 속성으로 정의하면 매우 쉽게 object content 를 저장하거나 읽을 수 있다. 하지만, 이렇게 하면 content 전체가 메모리에 올라가는 문제가 발생한다. content 크기가 1Mega byte 면 문제 없겠지만 1Giga byte, 1Tera byte 면 메모리 부족 오류가 발생할 수 밖에 없다.

```go
type Object struct {
	id        string
	container *Container
	content   []byte
}
```

이 문제는 Stream 방식으로 해결할 수 있다. go 의 io 패키지는 io stream 처리를 위한 훌륭한 인터페이스와 구현체들을 제공한다. go io 패키지는 io.Reader, io.Writer, io.Closer, io.Seeker 인터페이스를 제공한다. SOLID [Interface segregation principle](https://en.wikipedia.org/wiki/Interface_segregation_principle) - 인터페이스 분리 원칙에 충실하게 디자인 되어 있다. Read, Write, Close, Seek 기능을 개별 인터페이스로 분리해서 클라이언트 에서 필요한 인터페이스만을 사용할 수 있도록 한다. bytes.Buffer 패키지는 데이터를 버퍼에 쓰거나 읽을 수 있는 기능을 제공하기 위해서 io.Reader, io.Write 를 구현한다.  bytes.Reader 는 버퍼에 있는 데이터를 io.Reader 인터페이스로 제공한다. hash 패키지는 io.Write 를 구현함으로서 스트림 데이터의 hash 값을 `io.Copy(hash, reader)` 한 줄로 매우 쉽게 구할 수 있도록 해준다.

object 는 content 를 쓰거나 읽을 수 있어야 한다. 이를 위해서, io.Reader, io.Writer 인터페이스를 구현하겠다. 여기서, 하나 고민되는 점이 있다. object 정보와 ( 이하 object meta ), object content 가 저장되는 storage 로 file system, S2, DBMS 등 다양한 infrastructure 를 지원하고 싶다. 다르게 얘기하면, SOLID [Dependency inversion principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle) 에 따라 object meta, content 를 infrastructure 구현체가 아닌 추상화된 interface 에 의존 하고 싶다. 이미, go 에는 io.Reader, io.Writer 추상 인터페이스가 있다. 따라서, 인터페이스 구현 객체 (infrastructure 구현체)를 object 객체 생성시 주입하여 io.Reader, io.Writer 의 역할을 위임 시킬 수 있다. 이제, 아래와 같이 object type 을 정의 할 수 있다. io.Reader, io.Writer 를 go interface embedding 을 사용해서 속성으로 추가한다. 이제, Object.Read(p []byte), Object.Write(p []byte) 함수를 외부에서 사용할 수 있다.

```go
type Object struct {
	id        string
	container *Container
	length    int
	io.Reader
	io.Writer
}
```

우리가 정의한 object 는 편의상 동시에 읽고, 쓸 수는 없도록 하겠다. 이렇게 하기 위해서 읽기 전용 object 생성자, 쓰기 전용 object 생성자를 만들고, 읽기 object 인 경우, Write 함수 호출시 에러를 발생시키고, 마찬가지로 쓰기 object 에 Read 호출시 에러를 발생시키도록 하겠다. OpenObjectForRead 생성자는 object meta 정보와 함께 object Read() 를 위임할 io.Reader 를 주입받는다. 마찬가지로 OpenObjectForWrite 생성자는 object meta 정보와 함께 io.Writer를 주입받는다.

읽기 object 에 Write 에러를 내기 위해서 io.Reader 를 embedding 하지 않고 reader io.Reader 로 선언하고 직접 Read(p []byte) 를 구현해서 io.Reader 를 구현하는 것으로 변경한다. io.Writer 도 마찬가지이고, 직접 구현한 Read, Write 함수에서 오류 확인을 하고 각각 주입된 reader, writer 를 호출한다.

[domain/object.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/object.go)

```go
type Object struct {
   id        string
   container *Container
   length    int
   reader    io.Reader
   writer    io.Writer
}

func OpenObjectForRead(id string, container *Container, length int, reader io.Reader) *Object {
   return &Object{id: id, container: container, length: length, reader: reader}
}
func OpenObjectForWrite(id string, container *Container, writer io.Writer) *Object {
   return &Object{id: id, container: container, writer: writer}
}

var NotOpenForReadError = errors.New("object is not open for read")
var NotOpenForWriteError = errors.New("object is not open for write")

func (o *Object) Len() int {
	return o.length
}
func (o *Object) Read(p []byte) (n int, err error) {
	if o.reader == nil {
		return 0, NotOpenForReadError
	}
	return o.reader.Read(p)
}
func (o *Object) Write(p []byte) (n int, err error) {
	if o.writer == nil {
		return 0, NotOpenForWriteError
	}
	n, err = o.writer.Write(p)
	o.length += n
	return
}
```

infrastructure 구현체에 따라 Close 를 명시적으로 해야만 하는 경우들이 있다. 네트웍 통신, 파일 시스템 등이 대표적인 경우이다. 이를 위해서 object 에 io.Closer 를 구현한다.

```go
func (o *Object) Close() error {
   if o.reader != nil {
      if closer, ok := o.reader.(io.Closer); ok {
         return closer.Close()
      }
   }
   if o.writer != nil {
      if closer, ok := o.writer.(io.Closer); ok {
         return closer.Close()
      }
   }
   return nil
}
```

 이제 object 모델 정의가 완성됐다. object 모델은 account, container 보다 훨씬 복잡하다. 또, content 를 읽고, 쓰고, 닫는 기능도 있다. 테스트가 절실하다. 테스트 코드를 통해서 go io 패키지가 얼마나 훌륭한지 알게 되고, 우리가 정의한 object 모델이 적절한지 판단할 수 있을 것이다.

첫번째 테스트 - "read object with object.Read" 는 object.Read() 함수를 직접 사용해서 데이터를 읽는다. bytes.Reader 를 생성하고 OpenObjectForRead() 생성자에 주입한다. object.Read(p) 로 len(p) = len(content) 만큼 데이터를 reader 에서 읽는다. 그 결과 bytes.Reader 가 제공하는 content 가 p 에 들어간다.

두번째 테스트 - "read object with io.Copy" 는 io.Copy(dst io.Writer, src io.Reader) 함수를 활용하는 테스트 이다. 첫번째 테스트와 비교해보면, Read(p) 호출자는 스스로 전체 데이터 길이만큼 p []byte 를 할당하는 한편  io.Copy 호출자는 읽을 데이터 크기를 관여하지 않고 데이터를 저장하는 것을 io.Writer 에게 위임한다. io.Writer 는 파일 시스템이 될 수 도 있고, 네트웍 Connection 이 될 수 도 있고, http response writer 가 될 수 도 있다. object 및 어플리케이션 로직 변경없이 io.Writer 인스턴스로 교체가능하다는 점에서 [Liskov substitution principle](https://en.wikipedia.org/wiki/Liskov_substitution_principle) 을 따르고 있다.

[domain/object_test.go](https://github.com/reuben-baek/clean-go-application/blob/v0_1/domain-definition/domain/object_test.go)

```go
func TestObject(t *testing.T) {
	reuben := domain.NewAccount("reuben")
	document := domain.NewContainer("document", reuben)

	t.Run("read object with object.Read", func(t *testing.T) {
		content := []byte("hello world")
		reader := bytes.NewReader(content)
		object := domain.OpenObjectForRead("hello.txt", document, len(content), reader)

		p := make([]byte, object.Len())
		_, err := object.Read(p)
		assert.Nil(t, err)
		assert.Equal(t, content, p)

		err = object.Close()
		assert.Nil(t, err)
	})
	t.Run("read object with io.Copy", func(t *testing.T) {
		content := []byte("hello world")
		reader := bytes.NewReader(content)
		object := domain.OpenObjectForRead("hello.txt", document, len(content), reader)

		buffer := bytes.NewBuffer(nil)
		_, err := io.Copy(buffer, object)
		assert.Nil(t, err)
		assert.Equal(t, content, buffer.Bytes())

		err = object.Close()
		assert.Nil(t, err)
	})
  ...
}
```

세번째, 네번째 테스트는 io.Writer 로 주입된 writer 에 데이터를 Write 한다. object.Write, io.Copy 를 사용하는 전형적인 코드를 볼 수 있다.

```
func TestObject(t *testing.T) {
...
	t.Run("write object with object.Write", func(t *testing.T) {
		buffer := bytes.NewBuffer(nil)
		object := domain.OpenObjectForWrite("hello.txt", document, buffer)

		content := []byte("hello world")
		_, err := object.Write(content)
		assert.Nil(t, err)
		assert.Equal(t, len(content), object.Len())
		err = object.Close()
		assert.Nil(t, err)
		assert.Equal(t, content, buffer.Bytes())
	})
	t.Run("write object with io.Copy", func(t *testing.T) {
		buffer := bytes.NewBuffer(nil)
		object := domain.OpenObjectForWrite("hello.txt", document, buffer)

		content := []byte("hello world")
		reader := bytes.NewReader(content)

		_, err := io.Copy(object, reader)
		assert.Nil(t, err)
		assert.Equal(t, len(content), object.Len())
		err = object.Close()
		assert.Nil(t, err)
		assert.Equal(t, content, buffer.Bytes())
	})
}
```

테스트 코드로 부터 우리가 정의한 object 는 군더더기 없이 깔끔하게 데이터를 읽고, 쓰는 기능을 수행한다고 할 수 있다. object 를 이렇게 정의한 것은 [오브젝트 - 코드로 이해하는 객체지향 설계](https://wikibook.co.kr/object/) 의 자율적 객체 개념에서 힌트를 얻고, [io.File](https://golang.org/pkg/os/#File) 를 참조함으로서 가능했다. 다시 한번 강조하지만, go io 패키지는 매우 유연하면서도 강력한 추상 타입을 제공한다. io stream 처리를 할때에는 반드시 정확히 io 패키지를 이해하고, 어플리케이션 전반에서 추상 타입을 사용하고 인프라스트럭처 구현체를 주입해서 사용하라!!  어플리케이션 테스트가 가능해지고, 어플리케이션 stream 처리 코드가 간결해지는 것을 확인할 수 있을 것 이다. 

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

