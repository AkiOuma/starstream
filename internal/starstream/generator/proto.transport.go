package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/AkiOuma/starstream/internal/starstream/frame"
	"github.com/AkiOuma/starstream/internal/starstream/template"
)

func (g *Generator) buildProtoTransport(transport *frame.ProtoTransport) error {
	importPackages := bytes.NewBufferString("")
	for _, v := range transport.ImportPackages {
		importPackages.WriteString(fmt.Sprintf("import \"%s\";\n", v))
		if err := g.setThirdPartyProtoFIles(v); err != nil {
			return err
		}
	}
	content := fmt.Sprintf(
		template.ProtoTransportTemplate,
		transport.Package,
		transport.GoPackage,
		importPackages.String(),
		transport.Name,
		buildRpcMethods(transport.Method),
		buildRpcMessages(transport.Method),
	)
	return ioutil.WriteFile(transport.FilePath, []byte(content), 0644)
}

func buildRpcMethods(methods []*frame.RpcMethod) string {
	buf := bytes.NewBufferString("")
	for _, v := range methods {
		buf.WriteString(fmt.Sprintf(
			template.ProtoTransportRpcTemplate,
			v.Name,
			v.Request.Name,
			v.Reply.Name,
		))
	}
	return buf.String()
}

func buildRpcMessages(methods []*frame.RpcMethod) string {
	buf := bytes.NewBufferString("")
	for _, v := range methods {
		buf.WriteString(
			fmt.Sprintf(
				template.ProtoMessageTemplate,
				v.Request.Name,
				protoRpcMessageField(v.Request),
			),
		)
		buf.WriteString(
			fmt.Sprintf(
				template.ProtoMessageTemplate,
				v.Reply.Name,
				protoRpcMessageField(v.Reply),
			),
		)
	}
	return buf.String()
}

func protoRpcMessageField(proto *frame.Message) string {
	fields := bytes.NewBufferString("")
	for _, v := range proto.Field {
		fields.WriteString(fmt.Sprintf("\t%s\t%s = %d;\n", v.Type, v.Name, v.Id))
	}
	return fields.String()
}
