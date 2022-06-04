package template

var ValueObject = `package valueobject

import (
	%s
)

// Value Object
type %s struct {
	%s
}

// Querier
type %s struct {
	%s
}

`
