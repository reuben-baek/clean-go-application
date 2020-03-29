package domain_test

import (
	"bytes"
	"github.com/reuben-baek/clean-go-application/domain"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

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
