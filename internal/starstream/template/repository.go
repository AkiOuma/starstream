package template

var Repository = `package repository

import (
	%s
)

type %s interface {
	%s
}

// Finding Results
%s
`

var FindResult = `type %s struct {
	%s
}
`
