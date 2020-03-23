package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/infrastructure/inmemory"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAccountRouter_Put_Get(t *testing.T) {
	engine := gin.Default()
	accountRepository := inmemory.NewAccountRepository()
	accountApp := application.NewAccountApplication(accountRepository)
	accountRouter := NewAccountRouter(accountApp)
	rootRouter := newRootRouter(engine, accountRouter)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/reuben", nil)
	rootRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/reuben", nil)
	rootRouter.ServeHTTP(w, req)

	reuben := application.NewAccount("reuben")
	expected, _ := json.Marshal(reuben)
	assert.Equal(t, string(expected), strings.TrimSpace(w.Body.String()))
}
