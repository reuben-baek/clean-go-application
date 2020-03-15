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
}
