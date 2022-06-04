package frame

import (
	"log"
	"testing"
)

func TestNewRepository(t *testing.T) {
	def := getDef()
	info := &ServiceInfo{
		Name:        def.ServiceName,
		Destination: def.Destination,
	}
	entity := NewEntity(def.Entity[0], info)
	vo := NewValueObject(def.Entity[0], info)
	q := NewValueObjectQuerier(def.Entity[0])
	vo.Querier = q
	entity.ValueObject = vo
	repo := NewRepository(entity, info)
	log.Printf("%#v", repo)
}
