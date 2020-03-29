package domain_test

import (
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer(t *testing.T) {
	reuben := domain.NewAccount("reuben")
	document := domain.NewContainer("document", reuben)

	assert.Equal(t, "document", document.Id())
}
