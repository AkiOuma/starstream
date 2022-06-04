package frame

import (
	"path/filepath"
	"strings"
)

type Usecase struct {
	Name, ImplName       string
	ImportPath, FilePath string
	ImportPackages       []string
	Method               []*Method
	Entity               *Entity
}

func NewUsecase(e *Entity, serviceName, destination string) *Usecase {
	u := &Usecase{}
	u.Name = e.Name
	u.Entity = e
	u.ImplName = strings.ToLower(e.Name)
	u.ImportPath = filepath.Join(serviceName, "internal/usecase")
	u.FilePath = filepath.Join(destination, "internal/usecase", u.ImplName+".go")
	u.ImportPackages = make([]string, 0, 5)
	u.ImportPackages = append(u.ImportPackages,
		"context",
		e.ImportPath,
		e.ValueObject.ImportPath,
		e.Service.ImportPath,
		e.Repository.ImportPath,
	)
	for i, v := range []string{"Find", "Save", "Remove"} {
		methodName := v + e.Name + "s"
		method := NewMethod(methodName)
		method.AddArugs(&Param{Name: "ctx", DataType: "context.Context"})
		switch i {
		case 0:
			// find
			method.AddArugs(&Param{Name: "querier", DataType: "*valueobject." + e.ValueObject.Querier.Name})
			method.AddResults(&Param{DataType: "*repository." + e.Repository.FindResult.Name})
		case 1:
			// save
			method.AddArugs(&Param{Name: "data", DataType: "[]*entity." + e.Name})
			method.AddResults(&Param{DataType: "[]*entity." + e.Name})
		case 2:
			// remove
			method.AddArugs(&Param{Name: "querier", DataType: "*valueobject." + e.ValueObject.Querier.Name})
		default:
		}
		method.AddResults(&Param{DataType: "error"})
		u.Method = append(u.Method, method)
	}
	return u
}
