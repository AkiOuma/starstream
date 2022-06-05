package frame

import (
	"path/filepath"
	"strings"
)

type ProtoTransport struct {
	FilePath       string
	Package        string
	GoPackage      string
	ImportPackages []string
	Name           string
	Method         []*RpcMethod
	Message        []*Message
	Querier        *ProtoQuerier
	*ServiceInfo
}

type RpcMethod struct {
	Name    string
	Request *Message
	Reply   *Message
}

type Message struct {
	Name  string
	Field []*ProtoField
}

func NewProtoTransport(entity *Entity, proto *Proto) *ProtoTransport {
	transport := &ProtoTransport{}
	info := proto.ServiceInfo
	transport.ServiceInfo = info
	transport.Package = proto.Package
	transport.GoPackage = proto.GoPackage
	transport.Name = "Transport"
	transport.FilePath = filepath.Join(info.Destination, "api/proto", info.BriefServiceName, info.Type, info.Version, strings.ToLower(entity.Name)+".transport.proto")
	transport.ImportPackages = make([]string, 0, 1)
	transport.ImportPackages = append(transport.ImportPackages, filepath.Base(proto.FilePath))

	// methods
	transport.Message = make([]*Message, 0, len(entity.Usecase.Method)*2)
	transport.Method = make([]*RpcMethod, 0, len(entity.Usecase.Method))
	for i, v := range entity.Usecase.Method {
		method := &RpcMethod{}
		method.Name = v.Name
		request := &Message{Name: method.Name + "Request"}
		reply := &Message{Name: method.Name + "Reply"}
		method.Request = request
		method.Reply = reply
		request.Field = make([]*ProtoField, 0, 2)
		reply.Field = make([]*ProtoField, 0, 2)
		switch i {
		case 0:
			// find
			request.Field = append(request.Field, &ProtoField{
				Id:   1,
				Type: proto.Package + "." + proto.Name + "Querier",
				Name: "querier",
			})
			reply.Field = append(reply.Field, &ProtoField{
				Id:   1,
				Name: "count",
				Type: "int32",
			}, &ProtoField{
				Id:   2,
				Type: "repeated " + proto.Package + "." + proto.Name,
				Name: "data",
			})
		case 1:
			// save
			request.Field = append(request.Field, &ProtoField{
				Id:   1,
				Type: "repeated " + proto.Package + "." + proto.Name,
				Name: "data",
			})
			reply.Field = append(reply.Field, &ProtoField{
				Id:   1,
				Type: "repeated " + proto.Package + "." + proto.Name,
				Name: "data",
			})
		case 2:
			// remove
			request.Field = append(request.Field, &ProtoField{
				Id:   1,
				Type: proto.Package + "." + proto.Name + "Querier",
				Name: "querier",
			})
		}
		transport.Method = append(transport.Method, method)
		transport.Message = append(transport.Message, request, reply)
	}
	return transport
}
