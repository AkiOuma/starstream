package frame

import (
	"log"
	"testing"
)

func TestNewProtoQuerier(t *testing.T) {
	def := getDef()
	f := NewFrame(def)
	e := NewProtoQuerier(f.Proto[0])
	log.Printf("%#v", e)
}
