package frame

import "github.com/AkiOuma/starstream/internal/starstream/definition"

type Frame struct {
	ServiceName string
	Destination string
	Entity      []*Entity
	Proto       []*Proto
}

func NewFrame(def *definition.Definition) *Frame {
	f := &Frame{}
	f.ServiceName = def.ServiceName
	f.Destination = def.Destination
	f.Entity = make([]*Entity, 0, len(def.Entity))
	f.Proto = make([]*Proto, 0, len(def.Entity))
	for _, v := range def.Entity {
		entity := NewEntity(v, f.ServiceName, f.Destination)
		vo := NewValueObject(v, f.ServiceName, f.Destination)
		querier := NewValueObjectQuerier(v)
		vo.SetQuerier(querier)
		entity.SetValueObject(vo)
		repo := NewRepository(entity, f.ServiceName, f.Destination)
		entity.SetRepository(repo)
		service := NewService(entity, f.ServiceName, f.Destination)
		entity.Service = service
		usecase := NewUsecase(entity, f.ServiceName, f.Destination)
		entity.Usecase = usecase

		proto := NewProto(v, def.ServiceName, def.Destination)
		protoQuerier := NewProtoQuerier(v)
		proto.Querier = protoQuerier

		f.Entity = append(f.Entity, entity)
		f.Proto = append(f.Proto, proto)
	}
	return f
}
