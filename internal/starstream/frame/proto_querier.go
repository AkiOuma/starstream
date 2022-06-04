package frame

import (
	"github.com/AkiOuma/starstream/internal/starstream/definition"
)

type ProtoQuerier struct {
	Name  string
	Field []*ProtoQuerierField
}

type ProtoQuerierField struct {
	Id   int
	Name string
	Type string
}

func NewProtoQuerier(def *definition.Entity) *ProtoQuerier {
	q := &ProtoQuerier{}
	q.Name = GetPublicName(def.Name) + "Querier"
	q.Field = make([]*ProtoQuerierField, 0, len(def.Field))
	i := 0
	for _, v := range def.Field {
		switch v.Type {
		case "int":
			fieldType := ConvertProtoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name,
				Type: "repeated " + fieldType.TypeName,
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: v.Name + "Lower",
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 3,
				Name: v.Name + "Upper",
				Type: fieldType.TypeName,
			})
			i += 3
		case "string":
			fieldType := ConvertProtoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: "repeated " + v.Name,
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: "Search" + v.Name,
				Type: fieldType.TypeName,
			})
			i += 2
		case "time":
			fieldType := ConvertProtoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name + "Lower",
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: v.Name + "Upper",
				Type: fieldType.TypeName,
			})
			i += 2
		case "float":
			fieldType := ConvertProtoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name + "Lower",
				Type: fieldType.TypeName,
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: v.Name + "Upper",
				Type: fieldType.TypeName,
			})
			i += 2
		case "bool":
			fieldType := ConvertProtoType(v.Type)
			if fieldType == nil {
				continue
			}
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name,
				Type: fieldType.TypeName,
			})
		}
	}
	return q
}
