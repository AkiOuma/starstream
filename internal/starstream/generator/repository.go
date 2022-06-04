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

func (g *Generator) buildRepository(repo *frame.Repository) error {
	folder := filepath.Dir(repo.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	methods := bytes.NewBufferString("")
	for _, v := range repo.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("\"%s\"\n", v))
	}
	for _, v := range repo.Method {
		methods.WriteString(fmt.Sprintf(
			"\t%s(%s)(%s)\n",
			v.Name,
			buildParam(v.Args...),
			buildParam(v.Results...),
		))
	}
	content := fmt.Sprintf(
		template.Repository,
		importPackages.String(),
		repo.Name,
		methods.String(),
		buildFindingResult(repo.FindResult),
	)
	data, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(repo.FilePath, data, 0644)
}

func buildFindingResult(source *frame.FindResult) string {
	fields := bytes.NewBufferString("")
	for _, v := range source.Field {
		fields.WriteString(fmt.Sprintf("\t%s %s\n", v.Name, v.DataType))
	}
	return fmt.Sprintf(
		template.FindResult,
		source.Name,
		fields.String(),
	)
}
