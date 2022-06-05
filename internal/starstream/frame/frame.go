package frame

import (
	"strings"

	"github.com/AkiOuma/starstream/internal/starstream/definition"
)

type Frame struct {
	*ServiceInfo
	Entity []*Entity
	Proto  []*Proto
}

type ServiceInfo struct {
	Name             string
	Destination      string
	Version          string
	Type             string
	BriefServiceName string
}

func NewFrame(def *definition.Definition) *Frame {
	f := &Frame{}
	briefServiceName := ""
	list := strings.Split(def.ServiceName, "/")
	if len(list) > 0 {
		briefServiceName = list[len(list)-1]
	}
	f.ServiceInfo = &ServiceInfo{
		Name:             def.ServiceName,
		Destination:      def.Destination,
		Version:          def.ApiVersion,
		Type:             def.Type,
		BriefServiceName: briefServiceName,
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
		transport := NewTransport(entity, usecase, vo, f.ServiceInfo)
		entity.Transport = transport

		proto := NewProto(entity, f.ServiceInfo)
		protoQuerier := NewProtoQuerier(proto)
		proto.Querier = protoQuerier
		protoTransport := NewProtoTransport(entity, proto)
		proto.Transport = protoTransport

		f.Entity = append(f.Entity, entity)
		f.Proto = append(f.Proto, proto)
	}
	return f
}
