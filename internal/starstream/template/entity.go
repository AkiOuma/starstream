package template

var Entity = `package entity

import (
	%s
)

// Entity
type %s struct {
	%s
}

// Constructor
%s

// Getter 
%s

`

var NewEntity = `func New%s(root %s, vo *valueobject.%s) *%s {
	return &%s{
		%s
	}
}
`

var Getter = `func (e *%s) %s %s {
	return e.%s
}
`
