package frame

import (
	"log"
	"testing"

	"github.com/AkiOuma/starstream/internal/starstream/definition"
)

func TestPublicName(t *testing.T) {
	source, target := "user", "User"
	if GetPublicName(source) != target {
		t.Error("GetPublicName do not pass")
	}
}

func getDef() *definition.Definition {
	dir := "/root/tool/go/flame/starstream/example/demo01/definition.yaml"
	d, err := definition.ReadDefinition(dir)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
