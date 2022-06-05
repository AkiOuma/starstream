package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AkiOuma/starstream/internal/starstream/frame"
	"github.com/AkiOuma/starstream/internal/starstream/template"
)

func (g *Generator) buildTransport(transport *frame.Transport) error {
	return nil
}

func (g *Generator) buildTransportHelper(helper *frame.TransportHelper) error {
	// mkdir -p
	folder := filepath.Dir(helper.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	// fill template
	content := fmt.Sprintf(
		template.TransportHelperTemplate,
		helper.ImportPackages[0],
		helper.ImportPackages[1],
		helper.ImportPackages[2],
		proto2Entity(helper.Converter[1]),
		entity2Proto(helper.Converter[0]),
		protoQuerier2VoQuerier(helper.Converter[2]),
	)
	// genenrate file
	formatted, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(helper.FilePath, formatted, 0644)
}

func proto2Entity(convert *frame.TransportConverter) string {
	convertField := bytes.NewBufferString("")
	for _, v := range convert.Field {
		if v.Root {
			continue
		}
		convertField.WriteString(fmt.Sprintf("%s: %s,\n", v.FieldName, v.Call))
	}
	return fmt.Sprintf(
		template.Proto2EntityTemplate,
		convert.Name,
		convert.Source,
		convert.Target,
		convert.Target,
		convert.EntityName,
		convert.Root.Call,
		convert.EntityName,
		convertField.String(),
	)
}

func entity2Proto(convert *frame.TransportConverter) string {
	convertField := bytes.NewBufferString("")
	for _, v := range convert.Field {
		if v.Root {
			continue
		}
		convertField.WriteString(fmt.Sprintf("%s: %s,\n", v.FieldName, v.Call))
	}
	return fmt.Sprintf(
		template.Entity2ProtoTemplate,
		convert.Name,
		convert.Source,
		convert.Target,
		convert.Target,
		convert.EntityName,
		convertField.String(),
	)
}

func protoQuerier2VoQuerier(convert *frame.TransportConverter) string {
	convertField := bytes.NewBufferString("")
	for _, v := range convert.Field {
		convertField.WriteString(fmt.Sprintf("%s: %s,\n", v.FieldName, v.Call))
	}
	return fmt.Sprintf(
		template.ConvertQuerierTemplate,
		convert.Name,
		convert.Source,
		convert.Target,
		convert.VOName,
		convertField.String(),
	)
}
