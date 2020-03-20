package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/reuben-baek/clean-go-application/infrastructure/inmemory"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestAccountRouter(t *testing.T) {
	accountRepository := inmemory.NewAccountRepository()
	accountApp := application.NewAccountApplication(accountRepository)
	accountRouter := NewAccountRouter(accountApp)

	reuben := domain.NewAccount("reuben")
	accountRepository.Save(reuben)

	router := RootRouter(gin.Default())
	router.Handle(accountRouter)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/reuben", nil)
	router.ServeHTTP(w, req)

	expected := fmt.Sprintf("Hello, %+v", reuben)
	assert.Equal(t, expected, w.Body.String())
}
