package frame

import (
	"log"
	"testing"
)

func TestNewValueObject(t *testing.T) {
	def := getDef()
	for _, v := range def.Entity {
		e := NewValueObject(v, def.ServiceName, def.Destination)
		log.Printf("%v", e)
	}
}
