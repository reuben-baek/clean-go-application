package domain

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestNotFoundError_Error(t *testing.T) {
	const message = "cannot find reuben in account repository"
	pipeError := io.ErrClosedPipe
	err := NewNotFoundError(message, pipeError)
	expected := "NotFoundError: " + message + "; " + pipeError.Error()
	assert.Equal(t, expected, err.Error())
}

func TestNotFoundError_Unwrap(t *testing.T) {
	const message = "cannot find reuben in account repository"
	pipeError := io.ErrClosedPipe
	err := NewNotFoundError(message, pipeError)
	assert.Equal(t, pipeError, err.Unwrap())
}
