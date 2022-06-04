package generator

import (
	"log"
	"testing"

	"github.com/AkiOuma/starstream/internal/starstream/definition"
	"github.com/AkiOuma/starstream/internal/starstream/frame"
)

func TestGenerator(t *testing.T) {
	g := NewGenerator(frame.NewFrame(getDef()))
	g.Launch()
}

func getDef() *definition.Definition {
	dir := "/root/tool/go/flame/starstream/test/definition.yaml"
	d, err := definition.ReadDefinition(dir)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
