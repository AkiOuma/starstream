package template

var Service = `package service
import (
	%s
)

type %s interface {}

var _ %s = (*%s)(nil)

type %s struct {
	repo repository.%s
}
`
