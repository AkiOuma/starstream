package main

import (
	"flag"
	"log"

	"github.com/AkiOuma/starstream/internal/starstream/definition"
	"github.com/AkiOuma/starstream/internal/starstream/frame"
	"github.com/AkiOuma/starstream/internal/starstream/generator"
)

func main() {
	defDir := flag.String("defdir", "./definition.yaml", "path of configuration file, default value is './definition.yaml'")
	flag.Parse()
	def, err := definition.ReadDefinition(*defDir)
	if err != nil {
		log.Fatal(err.Error())
	}
	frame := frame.NewFrame(def)
	generator := generator.NewGenerator(frame)

	generator.Launch()
	// template.Bulidcfg("/root/tool/go/flame/starstream/internal/starstream/template/configmap.yaml", "/root/tool/go/flame/starstream/internal/starstream/template/configmap1.yaml")
}
