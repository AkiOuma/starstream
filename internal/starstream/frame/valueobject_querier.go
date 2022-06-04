package frame

import (
	"github.com/AkiOuma/starstream/internal/starstream/definition"
)

type ValueObjectQuerier struct {
	Name  string
	Field []*ValueObjectQuerierField
}

type ValueObjectQuerierField struct {
	Id   int
	Name string
	Type string
}

func NewValueObjectQuerier(def *definition.Entity) *ValueObjectQuerier {
	q := &ValueObjectQuerier{}
	q.Name = GetPublicName(def.Name) + "Querier"
	q.Field = make([]*ValueObjectQuerierField, 0, len(def.Field))
	i := 0
	for _, v := range def.Field {
		switch v.Type {
		case "int":
			fieldType := ConvertGoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 1,
				Name: GetPublicName(v.Name),
				Type: "[]" + fieldType.TypeName,
			})
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 2,
				Name: GetPublicName(v.Name) + "Lower",
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 3,
				Name: GetPublicName(v.Name) + "Upper",
				Type: fieldType.TypeName,
			})
			i += 3
		case "string":
			fieldType := ConvertGoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 1,
				Name: GetPublicName(v.Name),
				Type: "[]" + fieldType.TypeName,
			})
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 2,
				Name: "Search" + GetPublicName(v.Name),
				Type: fieldType.TypeName,
			})
			i += 2
		case "time":
			fieldType := ConvertGoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 1,
				Name: GetPublicName(v.Name) + "Lower",
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 2,
				Name: GetPublicName(v.Name) + "Upper",
				Type: fieldType.TypeName,
			})
			i += 2
		case "float":
			fieldType := ConvertGoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 1,
				Name: GetPublicName(v.Name) + "Lower",
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 2,
				Name: GetPublicName(v.Name) + "Upper",
				Type: fieldType.TypeName,
			})
			i += 2
		case "bool":
			fieldType := ConvertGoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ValueObjectQuerierField{
				Id:   i + 1,
				Name: GetPublicName(v.Name),
				Type: fieldType.TypeName,
			})
		}
	}
	return q
}
