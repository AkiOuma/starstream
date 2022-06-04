package frame

import (
	"log"
	"testing"
)

func TestNewProtoEntity(t *testing.T) {
	def := getDef()
	for _, v := range def.Entity {
		e := NewProto(v, def.ServiceName, def.Destination)
		log.Printf("%v", e)
	}
}
