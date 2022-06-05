package frame

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Proto struct {
	FilePath           string
	CompliedFolderPath string
	Package            string
	GoPackage          string
	ImportPackages     []string
	Name               string
	Field              []*ProtoField
	Querier            *ProtoQuerier
	*ServiceInfo
}

type ProtoField struct {
	Id   int
	Root bool
	Name string
	Type string
	// Getter string
}

func NewProto(entity *Entity, info *ServiceInfo) *Proto {
	importSet := make(map[string]bool)
	proto := &Proto{}
	proto.ServiceInfo = info
	proto.FilePath = filepath.Join(info.Destination, "api/proto", info.BriefServiceName, info.Type, info.Version, strings.ToLower(entity.Name)+".proto")
	proto.CompliedFolderPath = filepath.Join(info.Destination, "api/goapi", info.BriefServiceName, info.Type, info.Version)
	proto.Name = GetPublicName(entity.Name)
	proto.Field = make([]*ProtoField, 0, len(entity.Field))
	proto.Package = fmt.Sprintf("%s.%s.%s", info.BriefServiceName, info.Type, info.Version)
	proto.GoPackage = fmt.Sprintf("%s/%s/%s", info.BriefServiceName, info.Type, info.Version)
	for i, v := range entity.Field {
		fieldType := ConvertProtoType(v.Type)
		if fieldType == nil {
			continue
		}
		proto.Field = append(proto.Field, &ProtoField{
			Id:   i + 1,
			Name: v.PrivateName,
			Root: v.Root,
			Type: fieldType.TypeName,
			// Getter: "Get" + v.PublicName + "()",
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
