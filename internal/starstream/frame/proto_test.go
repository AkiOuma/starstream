package frame

import (
	"log"
	"testing"
)

func TestNewProtoEntity(t *testing.T) {
	def := getDef()
	f := NewFrame(def)
	for _, v := range f.Entity {
		e := NewProto(v, f.ServiceInfo)
		log.Printf("%#v", e)
	}
}
