# Simple Web Object Storage

account, container, object 를 서비스하는 간단한 web object storage 를 만들어 보자. account 는 container 들을 가질 수 있다. container 는 하부 container 를 가질 수 없고 object 들을 가진다. container name 은 '/' 를 포함하면 안된다. object name 은 '/' 를 포함한 path 형태로 표현된다.

```
/:account/:container/:object

ex) /reuben/documents/hello.txt
- account : reuben, container : documents, object : hello.txt

ex) /reuben/documents/research/go-research.md
- account : reuben, container : documents, object : research/go-research.md
```

account, container, object 에 대해 각각 create, read, update, delete 하는 rest api 를 제공한다.



## account

```
PUT /:account
- 신규 account를 생성한다.
GET /:account
- account 메타정보와 account의 container 목록을 제공한다.
POST /:account
- account 메타정보를 수정한다.
DELETE /:account
- account를 삭제한다.
```



## container

```
PUT /:account/:container
- account 에 새로운 container 를 생성한다.
GET /:account/:container
- container 메타정보와 container 의 object 목록을 제공한다.
POST /:account/:container
- container 의 메타정보를 수정한다.
DELETE /:account/:container
- container를 삭제한다.
```



## object

```
PUT /:account/:container/:object
- container 에 object를 생성 또는 교체한다.
GET /:account/:container/:object
- object 메타정보와 object data를 제공한다.
POST /:account/:container/:object
- object 메타정보를 수정한다.
DELETE /:account/:container/:object
- object를 삭제한다.
```



account, container, object 의 메타정보는 필요할때 정의한다.