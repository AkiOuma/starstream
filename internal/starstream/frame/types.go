package frame

type dataType struct {
	TypeName      string
	ImportPackage string
}

var goTypeMapper = map[string]*dataType{
	"int":    {TypeName: "int", ImportPackage: ""},
	"string": {TypeName: "string", ImportPackage: ""},
	"float":  {TypeName: "float64", ImportPackage: ""},
	"bool":   {TypeName: "bool", ImportPackage: ""},
	"time":   {TypeName: "time.Time", ImportPackage: "time"},
}

var protoTypeMapper = map[string]*dataType{
	"int":       {TypeName: "int32", ImportPackage: ""},
	"string":    {TypeName: "string", ImportPackage: ""},
	"float64":   {TypeName: "double", ImportPackage: ""},
	"bool":      {TypeName: "bool", ImportPackage: ""},
	"time.Time": {TypeName: "google.protobuf.Timestamp", ImportPackage: "google/protobuf/timestamp.proto"},
}

func ConvertGoType(origin string) *dataType {
	return goTypeMapper[origin]
}

func ConvertProtoType(origin string) *dataType {
	return protoTypeMapper[origin]
}
