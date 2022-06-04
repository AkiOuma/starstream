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

func (g *Generator) buildUsecase(uc *frame.Usecase) error {
	folder := filepath.Dir(uc.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	for _, v := range uc.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("\"%s\"\n", v))
	}
	content := fmt.Sprintf(
		template.Usecase,
		importPackages.String(),
		buildUsecaseInterface(uc),
		buildUsecaseImpl(uc),
		// interface validate
		uc.Name,
		uc.ImplName,
		buildUsecaseConstructor(uc),
		buildUsecaseImplMethods(uc),
	)
	data, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(uc.FilePath, data, 0644)
}

func buildUsecaseInterface(uc *frame.Usecase) string {
	return fmt.Sprintf(
		template.UsecaseInterface,
		uc.Name,
		buildUsecaseInterfaceMethod(uc),
	)
}

func buildUsecaseInterfaceMethod(uc *frame.Usecase) string {
	buf := bytes.NewBufferString("")
	for _, v := range uc.Method {
		buf.WriteString(fmt.Sprintf(
			"%s(%s)(%s)\n",
			v.Name,
			buildParam(v.Args...),
			buildParam(v.Results...),
		))
	}
	return buf.String()
}

func buildUsecaseImpl(uc *frame.Usecase) string {
	return fmt.Sprintf(
		template.UsecaseImplement,
		uc.ImplName,
		uc.Name,
		uc.Name,
	)
}

func buildUsecaseConstructor(uc *frame.Usecase) string {
	return fmt.Sprintf(
		template.UsecaseConstructor,
		uc.Name,
		uc.Name,
		uc.Name,
		uc.Name,
		uc.ImplName,
	)
}

func buildUsecaseImplMethods(uc *frame.Usecase) string {
	buf := bytes.NewBufferString("")
	content := func(i int) string {
		result := ""
		switch i {
		case 0:
			result = fmt.Sprintf("\treturn u.repo.%s(ctx, querier)", uc.Method[i].Name)
		case 1:
			result = fmt.Sprintf("\treturn u.repo.%s(ctx, data)", uc.Method[i].Name)
		case 2:
			result = fmt.Sprintf(
				template.UsecaseRemoveTemplate,
				uc.Name,
				uc.Entity.RootField.Type,
				uc.Entity.RootField.Getter,
				uc.Name,
			)
		}
		return result
	}
	for i, v := range uc.Method {
		buf.WriteString(fmt.Sprintf(
			template.UsecaseImplMethod,
			uc.ImplName,
			v.Name,
			buildParam(v.Args...),
			buildParam(v.Results...),
			content(i),
		))
	}
	return buf.String()
}
