package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
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
	if err := os.MkdirAll(proto.CompliedFolderPath, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	for _, v := range proto.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("import \"%s\";\n", v))
		if err := g.setThirdPartyProtoFIles(v); err != nil {
			return err
		}
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

func (g *Generator) setThirdPartyProtoFIles(moduleName string) error {
	dest := filepath.Join(
		g.frame.Destination,
		"third_party",
		filepath.Dir(moduleName),
	)
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		log.Fatalf("ERROR: Generate %s failed because: %v", dest, err)
	}
	var protoContent string
	switch moduleName {
	case "google/protobuf/timestamp.proto":
		protoContent = template.GoogleTimeStamp
	}
	return ioutil.WriteFile(
		filepath.Join(dest, filepath.Base(moduleName)),
		[]byte(protoContent),
		0644,
	)
}
