package frame

import "github.com/AkiOuma/starstream/internal/starstream/definition"

type Frame struct {
	*ServiceInfo
	Entity []*Entity
	Proto  []*Proto
}

type ServiceInfo struct {
	Name        string
	Destination string
	Version     string
	Type        string
}

func NewFrame(def *definition.Definition) *Frame {
	f := &Frame{}
	f.ServiceInfo = &ServiceInfo{
		Name:        def.ServiceName,
		Destination: def.Destination,
		Version:     def.ApiVersion,
		Type:        def.Type,
	}
	f.Entity = make([]*Entity, 0, len(def.Entity))
	f.Proto = make([]*Proto, 0, len(def.Entity))
	for _, v := range def.Entity {
		entity := NewEntity(v, f.ServiceInfo)
		vo := NewValueObject(v, f.ServiceInfo)
		querier := NewValueObjectQuerier(v)
		vo.SetQuerier(querier)
		entity.SetValueObject(vo)
		repo := NewRepository(entity, f.ServiceInfo)
		entity.SetRepository(repo)
		service := NewService(entity, f.ServiceInfo)
		entity.Service = service
		usecase := NewUsecase(entity, f.ServiceInfo)
		entity.Usecase = usecase

		proto := NewProto(entity, f.ServiceInfo)
		protoQuerier := NewProtoQuerier(v)
		proto.Querier = protoQuerier

		f.Entity = append(f.Entity, entity)
		f.Proto = append(f.Proto, proto)
	}
	return f
}
