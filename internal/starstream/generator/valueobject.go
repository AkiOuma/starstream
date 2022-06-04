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

func (g *Generator) buildValueObject(vo *frame.ValueObject) error {
	folder := filepath.Dir(vo.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	voFields := bytes.NewBufferString("")
	querierFields := bytes.NewBufferString("")
	for _, v := range vo.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("\"%s\"\n", v))
	}
	for _, v := range vo.Field {
		voFields.WriteString(fmt.Sprintf("%s %s\n", v.Name, v.Type))
	}
	for _, v := range vo.Querier.Field {
		querierFields.WriteString(fmt.Sprintf("%s %s\n", v.Name, v.Type))
	}
	content := fmt.Sprintf(
		template.ValueObject,
		importPackages.String(),
		vo.Name,
		voFields.String(),
		vo.Querier.Name,
		querierFields.String(),
	)
	data, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(vo.FilePath, data, 0644)
}
