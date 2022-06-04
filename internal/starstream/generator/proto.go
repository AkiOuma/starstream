package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AkiOuma/starstream/internal/starstream/frame"
	"github.com/AkiOuma/starstream/internal/starstream/template"
)

func (g *Generator) buildProto(proto *frame.Proto) error {
	folder := filepath.Dir(proto.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	for _, v := range proto.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("import \"%s\";\n", v))
	}
	content := fmt.Sprintf(
		template.ProtoTemplate,
		proto.Package,
		proto.GoPackage,
		importPackages.String(),
		proto.Name,
		protoEntityField(proto),
		proto.Querier.Name,
		protoQuerierField(proto.Querier),
	)
	return ioutil.WriteFile(proto.FilePath, []byte(content), 0644)
}

func protoEntityField(proto *frame.Proto) string {
	fields := bytes.NewBufferString("")
	for _, v := range proto.Field {
		fields.WriteString(fmt.Sprintf("\t%s\t%s = %d;\n", v.Type, v.Name, v.Id))
	}
	return fields.String()
}

func protoQuerierField(querier *frame.ProtoQuerier) string {
	fields := bytes.NewBufferString("")
	for _, v := range querier.Field {
		fields.WriteString(fmt.Sprintf("\t%s\t%s = %d;\n", v.Type, v.Name, v.Id))
	}
	return fields.String()
}

// var ProtoEntityTemplate = `syntax = "proto3";

// package %s;

// option go_package = "%s";

// %s

// message %s {
// %s
// }

// message %s {
// %s
// }
// `
