package frame

import (
	"log"
	"testing"
)

func TestNewValueObjectQuerier(t *testing.T) {
	def := getDef()
	for _, v := range def.Entity {
		e := NewValueObjectQuerier(v)
		log.Printf("%#v", e)
	}
}
