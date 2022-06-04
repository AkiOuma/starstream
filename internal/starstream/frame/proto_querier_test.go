package frame

import (
	"log"
	"testing"
)

func TestNewProtoQuerier(t *testing.T) {
	def := getDef()
	for _, v := range def.Entity {
		e := NewProtoQuerier(v)
		log.Printf("%#v", e)
	}
}
