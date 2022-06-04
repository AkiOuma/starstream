package frame

import (
	"log"
	"testing"
)

func TestNewUsecase(t *testing.T) {
	def := getDef()
	entity := NewEntity(def.Entity[0], def.ServiceName, def.Destination)
	vo := NewValueObject(def.Entity[0], def.ServiceName, def.Destination)
	q := NewValueObjectQuerier(def.Entity[0])
	vo.Querier = q
	entity.ValueObject = vo
	repo := NewRepository(entity, def.ServiceName, def.Destination)
	entity.Repository = repo
	service := NewService(entity, def.ServiceName, def.Destination)
	entity.Service = service
	usecase := NewUsecase(entity, def.ServiceName, def.Destination)
	log.Printf("%#v", usecase)
}
