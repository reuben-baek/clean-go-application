package application

import "github.com/reuben-baek/clean-go-application/domain"

type Object struct {
	Id string
}

func ObjectFrom(o domain.Object) Object {
	return Object{Id: o.Id()}
}
