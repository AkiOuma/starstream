package frame

import (
	"path/filepath"

	"github.com/AkiOuma/starstream/internal/starstream/definition"
)

type ValueObject struct {
	ImportPath, FilePath string
	ImportPackages       []string
	Name                 string
	Field                []*ValueObjectField
	Querier              *ValueObjectQuerier
}

type ValueObjectField struct {
	Id   int
	Name string
	Type string
}

func NewValueObject(def *definition.Entity, serviceName, destination string) *ValueObject {
	importSet := make(map[string]bool)
	vo := &ValueObject{}
	vo.ImportPath = filepath.Join(serviceName, "internal/domain/valueobject")
	vo.FilePath = filepath.Join(destination, "internal/domain/valueobject", def.Name+".go")
	vo.Name = GetPublicName(def.Name)
	vo.Field = make([]*ValueObjectField, 0, len(def.Field))
	vo.ImportPackages = make([]string, 0)
	for i, v := range def.Field {
		if v.Root {
			continue
		}
		fieldType := ConvertGoType(v.Type)
		if fieldType == nil {
			continue
		}
		vo.Field = append(vo.Field, &ValueObjectField{
			Id:   i + 1,
			Name: GetPublicName(v.Name),
			Type: fieldType.TypeName,
		})

		if len(fieldType.ImportPackage) > 0 && !importSet[fieldType.ImportPackage] {
			vo.ImportPackages = append(vo.ImportPackages, fieldType.ImportPackage)
			importSet[fieldType.ImportPackage] = true
		}
	}
	return vo
}

func (vo *ValueObject) SetQuerier(q *ValueObjectQuerier) {
	vo.Querier = q
}
