package inmemory

import (
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountRepository_Find_NotFound(t *testing.T) {
	var expectedError *domain.NotFoundError
	expectedError = domain.NewNotFoundError("cannot find reuben in inmemory.accountRepository", nil)
	accountRepository := NewAccountRepository()
	actual, err := accountRepository.Find("reuben")

	assert.Nil(t, nil, actual)
	assert.IsType(t, expectedError, err)
	assert.Equal(t, expectedError, err)
}

func TestAccountRepository_Save_Find(t *testing.T) {
	reuben := domain.NewAccount("reuben")
	accountRepository := NewAccountRepository()
	err := accountRepository.Save(reuben)
	assert.Nil(t, err)

	actual, err := accountRepository.Find(reuben.Id())
	assert.Nil(t, err)
	assert.Equal(t, reuben, actual)
}
