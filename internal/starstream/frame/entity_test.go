package frame

import (
	"log"
	"testing"
)

func TestNewGoEntity(t *testing.T) {
	def := getDef()
	for _, v := range def.Entity {
		e := NewEntity(v, &ServiceInfo{
			Name:        def.ServiceName,
			Destination: def.Destination,
		})
		log.Printf("%#v", e)
	}
}
