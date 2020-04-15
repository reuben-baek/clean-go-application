package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/reuben-baek/clean-go-application/application"
	"github.com/reuben-baek/clean-go-application/lib/webserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAccountRouter(t *testing.T) {
	accountApp := &mockAccountApplication{}
	accountRouter := NewAccountRouter(accountApp)
	rootRouter := webserver.NewRootRouter(gin.Default(), accountRouter)

	reuben := application.NewAccount("reuben")
	accountApp.On("Find", "reuben").Return(reuben, nil)
	accountApp.On("Save", reuben).Return(nil)

	t.Run("put /reuben", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("PUT", "/reuben", nil)
		rootRouter.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("get /reuben", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/reuben", nil)
		rootRouter.ServeHTTP(w, req)

		expected, _ := json.Marshal(reuben)
		assert.Equal(t, string(expected), strings.TrimSpace(w.Body.String()))
	})
}

type mockAccountApplication struct {
	mock.Mock
}

func (m *mockAccountApplication) Find(id string) (*application.Account, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*application.Account), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *mockAccountApplication) Save(account *application.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *mockAccountApplication) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}