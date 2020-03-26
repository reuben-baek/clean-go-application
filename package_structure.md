# Package Structure

패키지 및 디렉토리 구조는 DDD 구조와 같이 구성한다. domain 패키지에는 어플리케이션의 도메인 모델 - 엔터티, 서비스, 레파지토리들이 포함된다. infrastructure 패키지에는 도메인 레파지토리의 구현체들이 들어간다. application 패키지는 도메인 모델을 조합해서 어플리케이션의 입출력을 담당하는 구현체를 가진다. interfaces 패키지는 어플리케이션이 외부와 커뮤니케이션을 하기 위한 구체적인 코드를 포함한다. 좀더 자세한 사항은 뒤에서 단계별로 설명하겠다.

- interfaces
- application
- domain
- infrastructure

java, kotlin 에서는 한개의 파일에 한개의 class, interface 만을 정의하는게 보편적이다. 한 패키지에 정의된 class, interface 가 많아지면 자연스럽게 하부 패키지를 만들어 분류하게 된다. import 대상이 class, interface 이므로 코드에서 해당 type 을 바로 사용할 수 있고, 하부 패키지로 위치를 변경해도 import 의 패키지 경로만 변경되지 해당 코드가 변경되지 않는다.

```kotlin
import org.junit.Assert
import org.junit.Test

/**
 * @author reuben.baek
 */
class HelloTest {
    @Test
    fun hello() {
        val reuben = "reuben"
        val hello = "Hello, $reuben" 
        val expected = "Hello, reuben"
        Assert.assertEquals(expected, hello)
    }
}
```

 golang 에서는 이렇게 하기가 좀 어렵다. golang 에서는 외부에서 참조할때 package.Type ( ex, application.Account ) 의 형태를 사용한다. application 패키지에 파일이 많아져서 하위 패키지로 account 를 만들고 Account 를 이동시키면 코드에서 account.Account 로 참조하게 된다. 물론 alias 를 이용한다면 import 수정만으로 가능하지만, import 마다 alias 를 수동으로 해줘야 하는건 매우 귀찮고 권장할 만한 게 아니다. golang 의 package 를 java 의 class 로 생각하고 패키지 구성을 하는 것도 괜찮은데, 이럴 경우 디렉토리가 많아질 가능성이 있다. 그래도, 한 디렉토리에 많은 파일이 위치하는 것보다는 낫다.

```go
import (
   "fmt"
   "github.com/gin-gonic/gin"
   "github.com/reuben-baek/clean-go-application/application"
   "github.com/reuben-baek/clean-go-application/lib/webserver"
)

type AccountRouter struct {
   account *application.Account
}
```

 

```go
import (
   "fmt"
   "github.com/gin-gonic/gin"
   "github.com/reuben-baek/clean-go-application/application/account"
   "github.com/reuben-baek/clean-go-application/lib/webserver"
)

type AccountRouter struct {
   account *account.Account
}
```

golang 에서 패키지명을 고민하게 만드는 요인이 하나 더 있다. golang naming 규칙에 따르면 패키지 명은 소문자, _ 로 구성된다. '-' 를 사용하지 못한다. - 가 포함된 디렉토리명이 패키지이면 코드에는 _ 로 사용된다. 여기에 하나더 _ 를 사용하지 말것을 권고하고 있다. 예를들어 web-server 대신에 webserver 로 사용하라는 권고이다. httptest 와 같이.

여기에서는 하나의 파일에 관련있는 type 들을 모으고 ( IDE 가 좋아서 해당 코드로 이동하는데 전혀 문제가 없어서 ), 가능한 하부 패키지를 많이 만들지 않겠다. 물론, 독립적으로 명확히 패키지로 분리하는게 좋은 경우에는 만들것이다.

DDD layer 를 표현하는 네개 패키지 외에 두개의 디렉토리가 더 있다. lib 에는 DDD layer 와 무관한 helper, framework 등의 패키지가 들어가게 된다. cmd 에는 go main 패키지가 위치하게 된다.

- lib
- cmd