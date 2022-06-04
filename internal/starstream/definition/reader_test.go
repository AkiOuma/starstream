package definition

import (
	"log"
	"testing"
)

func TestReadDefinition(t *testing.T) {
	dir := "/root/tool/go/flame/starstream/test/definition.yaml"
	d, err := ReadDefinition(dir)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%#v", d)
}
