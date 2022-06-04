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

func (g *Generator) buildService(svc *frame.Service) error {
	folder := filepath.Dir(svc.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	for _, v := range svc.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("\"%s\"\n", v))
	}
	content := fmt.Sprintf(
		template.Service,
		importPackages.String(),
		svc.Name,
		svc.Name,
		svc.ImplName,
		svc.ImplName,
		svc.Repository.Name,
	)
	data, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(svc.FilePath, data, 0644)
}
