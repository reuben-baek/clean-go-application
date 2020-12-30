package domain

import (
	"errors"
	"io"
)

type Object interface {
	Id() string
	Len() int
	io.ReadWriteCloser
}

type object struct {
	id        string
	container Container
	length    int
	reader    io.Reader
	writer    io.Writer
}

func OpenObjectForRead(id string, container Container, length int, reader io.Reader) Object {
	return &object{id: id, container: container, length: length, reader: reader}
}
func OpenObjectForWrite(id string, container Container, writer io.Writer) Object {
	return &object{id: id, container: container, writer: writer}
}

var NotOpenForReadError = errors.New("object is not open for read")
var NotOpenForWriteError = errors.New("object is not open for write")

func (o *object) Id() string {
	return o.id
}
func (o *object) Len() int {
	return o.length
}

func (o *object) Read(p []byte) (n int, err error) {
	if o.reader == nil {
		return 0, NotOpenForReadError
	}
	return o.reader.Read(p)
}

func (o *object) Write(p []byte) (n int, err error) {
	if o.writer == nil {
		return 0, NotOpenForWriteError
	}
	n, err = o.writer.Write(p)
	o.length += n
	return
}

func (o *object) Close() error {
	if o.reader != nil {
		if closer, ok := o.reader.(io.Closer); ok {
			return closer.Close()
		}
	}
	if o.writer != nil {
		if closer, ok := o.writer.(io.Closer); ok {
			return closer.Close()
		}
	}
	return nil
}
