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

func (g *Generator) buildEntity(entity *frame.Entity) error {
	folder := filepath.Dir(entity.FilePath)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	importPackages := bytes.NewBufferString("")
	entityFields := bytes.NewBufferString("")
	for _, v := range entity.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("\"%s\"\n", v))
	}
	for _, v := range entity.Field {
		entityFields.WriteString(fmt.Sprintf("%s %s\n", v.PrivateName, v.Type))
	}
	content := fmt.Sprintf(
		template.Entity,
		importPackages.String(),
		entity.Name,
		entityFields.String(),
		entityConstructor(entity),
		entityGetter(entity),
	)
	data, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(entity.FilePath, data, 0644)
}

func entityConstructor(e *frame.Entity) string {
	fields := bytes.NewBufferString("")
	for _, v := range e.Field {
		if v.Root {
			fields.WriteString(fmt.Sprintf("%s: root,\n", v.PrivateName))
			continue
		}
		fields.WriteString(fmt.Sprintf("%s: vo.%s,\n", v.PrivateName, v.PublicName))
	}
	return fmt.Sprintf(
		template.NewEntity,
		e.Name,
		e.RootField.Type,
		e.Name,
		e.Name,
		e.Name,
		fields.String(),
	)
}

func entityGetter(e *frame.Entity) string {
	getter := bytes.NewBufferString("")
	for _, v := range e.Field {
		getter.WriteString(fmt.Sprintf(
			template.Getter,
			e.Name,
			v.Getter,
			v.Type,
			v.PrivateName,
		))
	}
	getter.WriteByte(10)
	return getter.String()
}
