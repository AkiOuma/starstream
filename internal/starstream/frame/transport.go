package frame

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Transport struct {
	Name                 string
	ImportPath, FilePath string
	ImportPackages       []string
	Method               []*Method
	Entity               *Entity
	ValueObject          *ValueObject
	Usecase              *Usecase
	Helper               *TransportHelper
	ServiceInfo          *ServiceInfo
}

type TransportHelper struct {
	Name                 string
	ImportPath, FilePath string
	ImportPackages       []string
	Converter            []*TransportConverter
}

type TransportConverter struct {
	Template       string
	Name           string
	EntityName     string
	VOName         string
	Source, Target string
	Field          []*ConvertField
	Root           *ConvertField
}

type ConvertField struct {
	Root            bool
	FieldName, Call string
}

func NewTransport(entity *Entity, usecase *Usecase, vo *ValueObject, info *ServiceInfo) *Transport {
	transport := &Transport{}
	transport.Name = "Transport"
	transport.ServiceInfo = info
	transport.FilePath = filepath.Join(info.Destination, "internal/interface/transport", strings.ToLower(entity.Name)+".go")
	transport.ImportPath = filepath.Join(info.Name, "internal/interface/transport")
	transport.Entity = entity
	transport.ValueObject = vo
	transport.Usecase = usecase
	helper := NewTranportHelper(transport)
	transport.Helper = helper
	return transport
}

func NewTranportHelper(transport *Transport) *TransportHelper {
	helper := &TransportHelper{}
	helper.FilePath = filepath.Join(transport.ServiceInfo.Destination, "internal/interface/transport/helper", strings.ToLower(transport.Entity.Name)+".go")
	helper.ImportPath = filepath.Join(transport.ServiceInfo.Name, "internal/interface/transport/helper")
	helper.ImportPackages = make([]string, 0, 3)
	helper.ImportPackages = append(
		helper.ImportPackages,
		filepath.Join(transport.ServiceInfo.Name, "api/goapi", transport.ServiceInfo.BriefServiceName, transport.ServiceInfo.Type, transport.ServiceInfo.Version),
	)
	helper.ImportPackages = append(helper.ImportPackages, transport.Entity.ImportPath)
	helper.ImportPackages = append(helper.ImportPackages, transport.Entity.ValueObject.ImportPath)

	helper.Converter = make([]*TransportConverter, 0, 3)

	// entity to proto
	converter := &TransportConverter{}
	converter.EntityName = transport.Entity.Name
	converter.Name = fmt.Sprintf("Entity%s2V1%s", transport.Entity.Name, transport.Entity.Name)
	converter.Source = "entity." + transport.Entity.Name
	converter.Target = "v1." + transport.Entity.Name
	for _, v := range transport.Entity.Field {
		field := &ConvertField{}
		field.FieldName = v.PublicName
		switch v.Type {
		case "int":
			field.Call = fmt.Sprintf("int32(v.%s)", v.Getter)
		case "time.Time":
			field.Call = fmt.Sprintf("timestamppb.New(v.%s)", v.Getter)
		default:
			field.Call = fmt.Sprintf("v.%s", v.Getter)
		}
		converter.Field = append(converter.Field, field)
	}
	helper.Converter = append(helper.Converter, converter)

	// proto to entity
	converter = &TransportConverter{}
	converter.EntityName = transport.Entity.Name
	converter.Name = fmt.Sprintf("V1%s2Entity%s", transport.Entity.Name, transport.Entity.Name)
	converter.Source = "v1." + transport.Entity.Name
	converter.Target = "entity." + transport.Entity.Name
	for _, v := range transport.Entity.Field {
		field := &ConvertField{}
		field.FieldName = v.PublicName
		switch v.Type {
		case "int":
			field.Call = fmt.Sprintf("int(v.%s)", v.Getter)
		case "time.Time":
			field.Call = fmt.Sprintf("v.%s.AsTime()", v.Getter)
		default:
			field.Call = fmt.Sprintf("v.%s", v.Getter)
		}
		field.Root = v.Root
		converter.Field = append(converter.Field, field)
		if field.Root {
			converter.Root = field
		}
	}
	helper.Converter = append(helper.Converter, converter)

	// proto querier to valueobject querier
	converter = &TransportConverter{}
	querierName := transport.ValueObject.Querier.Name
	converter.Name = fmt.Sprintf("V1%s2Entity%s", querierName, querierName)
	converter.Source = "v1." + querierName
	converter.Target = "valueobject." + querierName
	converter.VOName = transport.Entity.ValueObject.Querier.Name
	for _, v := range transport.ValueObject.Querier.Field {
		field := &ConvertField{}
		field.FieldName = v.Name
		switch v.Type {
		case "[]int":
			field.Call = fmt.Sprintf("bulkInt32ToInt(source.Get%s())", v.Name)
		case "int":
			field.Call = fmt.Sprintf("int(source.Get%s())", v.Name)
		case "time.Time":
			field.Call = fmt.Sprintf("source.Get%s().AsTime()", v.Name)
		default:
			field.Call = fmt.Sprintf("source.Get%s()", v.Name)
		}
		converter.Field = append(converter.Field, field)
	}
	helper.Converter = append(helper.Converter, converter)
	return helper
}
