package frame

type ProtoQuerier struct {
	Name  string
	Field []*ProtoQuerierField
}

type ProtoQuerierField struct {
	Id   int
	Name string
	Type string
}

func NewProtoQuerier(proto *Proto) *ProtoQuerier {
	q := &ProtoQuerier{}
	q.Name = GetPublicName(proto.Name) + "Querier"
	q.Field = make([]*ProtoQuerierField, 0, len(proto.Field))
	i := 0
	for _, v := range proto.Field {
		switch v.Type {
		case "int32":
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name,
				Type: "repeated int32",
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: v.Name + "Lower",
				Type: "int32",
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 3,
				Name: v.Name + "Upper",
				Type: "int32",
			})
			i += 3
		case "string":
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name,
				Type: "repeated string",
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: "Search" + v.Name,
				Type: "string",
			})
			i += 2
		case "google.protobuf.Timestamp":
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name + "Lower",
				Type: "google.protobuf.Timestamp",
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: v.Name + "Upper",
				Type: "google.protobuf.Timestamp",
			})
			i += 2
		case "double":
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 1,
				Name: v.Name + "Lower",
				Type: "double",
			})
			q.Field = append(q.Field, &ProtoQuerierField{
				Id:   i + 2,
				Name: v.Name + "Upper",
				Type: "double",
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
				Type: "int32",
			})
		}
	}
	return q
}
