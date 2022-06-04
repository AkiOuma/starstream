package frame

import (
	"log"
	"testing"
)

func TestNewUsecase(t *testing.T) {
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
	entity.Repository = repo
	service := NewService(entity, info)
	entity.Service = service
	usecase := NewUsecase(entity, info)
	log.Printf("%#v", usecase)
}
