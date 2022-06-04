package frame

import (
	"strings"
)

type Method struct {
	Name    string
	Args    []*Param
	Results []*Param
}

func NewMethod(name string) *Method {
	return &Method{
		Name:    name,
		Args:    make([]*Param, 0, 2),
		Results: make([]*Param, 0, 2),
	}
}

func (m *Method) AddArugs(args ...*Param) {
	m.Args = append(m.Args, args...)
}

func (m *Method) AddResults(args ...*Param) {
	m.Results = append(m.Results, args...)
}

type Param struct {
	Name     string
	DataType string
}

func GetPublicName(name string) string {
	if len(name) == 0 {
		return name
	}
	return strings.ToUpper(name[0:1]) + name[1:]
}
