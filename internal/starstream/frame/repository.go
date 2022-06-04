package frame

import (
	"path/filepath"
	"strings"
)

type Repository struct {
	Name                 string
	FilePath, ImportPath string
	ImportPackages       []string
	Method               []*Method
	FindResult           *FindResult
}

type FindResult struct {
	Name  string
	Field []*Param
}

func NewFindResult(entity *Entity) *FindResult {
	r := &FindResult{
		Name: "Find" + entity.Name + "sResult",
	}
	r.Field = make([]*Param, 0, 2)
	r.Field = append(r.Field, &Param{Name: "Count", DataType: "int"}, &Param{
		Name:     "Data",
		DataType: "[]*entity." + entity.Name,
	})
	return r
}

func NewRepository(entity *Entity, serviceName, destination string) *Repository {
	r := &Repository{}
	r.Name = entity.Name
	r.ImportPath = filepath.Join(serviceName, "internal/domain/repository")
	r.FilePath = filepath.Join(destination, "internal/domain/repository", strings.ToLower(entity.Name)+".go")
	r.Method = make([]*Method, 0, 3)
	r.ImportPackages = make([]string, 0)
	r.ImportPackages = append(r.ImportPackages, "context")
	if entity.ValueObject != nil {
		r.ImportPackages = append(r.ImportPackages, entity.ImportPath, entity.ValueObject.ImportPath)
	}
	r.FindResult = NewFindResult(entity)
	for i, v := range []string{"Find", "Save", "Remove"} {
		methodName := v + entity.Name + "s"
		method := NewMethod(methodName)
		method.AddArugs(&Param{Name: "ctx", DataType: "context.Context"})
		switch i {
		case 0:
			// find
			method.AddArugs(&Param{Name: "querier", DataType: "*valueobject." + entity.ValueObject.Querier.Name})
			method.AddResults(&Param{DataType: "*" + r.FindResult.Name})
		case 1:
			// save
			method.AddArugs(&Param{Name: "data", DataType: "[]*entity." + entity.Name})
			method.AddResults(&Param{DataType: "[]*entity." + entity.Name})
		case 2:
			// remove
			method.AddArugs(&Param{Name: "root", DataType: "[]" + entity.RootField.Type})
		default:
		}
		method.AddResults(&Param{DataType: "error"})
		r.Method = append(r.Method, method)
	}
	return r
}
