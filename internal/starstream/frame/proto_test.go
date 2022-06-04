package frame

import (
	"log"
	"testing"
)

func TestNewProtoEntity(t *testing.T) {
	def := getDef()
	for _, v := range def.Entity {
		entity := NewEntity(v, &ServiceInfo{
			Name:        def.ServiceName,
			Destination: def.Destination,
		})
		e := NewProto(entity, &ServiceInfo{
			Name:        def.ServiceName,
			Destination: def.Destination,
		})
		log.Printf("%v", e)
	}
}
