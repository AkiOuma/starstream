package frame

type Proto struct {
	FilePath       string
	Package        string
	GoPackage      string
	ImportPackages []string
	Name           string
	Field          []*ProtoField
	Querier        *ProtoQuerier
}

type ProtoField struct {
	Id     int
	Root   bool
	Name   string
	Type   string
	Getter string
}

func NewProto(def *Entity, info *ServiceInfo) *Proto {
	importSet := make(map[string]bool)
	proto := &Proto{}
	proto.Name = GetPublicName(def.Name)
	proto.Field = make([]*ProtoField, 0, len(def.Field))
	for i, v := range def.Field {
		fieldType := ConvertProtoType(v.Type)
		if fieldType == nil {
			continue
		}
		proto.Field = append(proto.Field, &ProtoField{
			Id:     i + 1,
			Name:   v.PrivateName,
			Root:   v.Root,
			Type:   fieldType.TypeName,
			Getter: "Get" + v.PublicName + "()",
		})
		if len(fieldType.ImportPackage) > 0 && !importSet[fieldType.ImportPackage] {
			proto.ImportPackages = append(proto.ImportPackages, fieldType.ImportPackage)
			importSet[fieldType.ImportPackage] = true
		}
	}
	return proto
}

func (vo *Proto) SetQuerier(q *ProtoQuerier) {
	vo.Querier = q
}
