package frame

import (
	"log"
	"testing"
)

func TestNewValueObject(t *testing.T) {
	def := getDef()
	info := &ServiceInfo{
		Name:        def.ServiceName,
		Destination: def.Destination,
	}
	for _, v := range def.Entity {
		e := NewValueObject(v, info)
		log.Printf("%v", e)
	}
}
