package template

var TransportTemplate = `package transport

import (
	"context"

	v1 "%s"
	"%s"
)

type Transport struct {
	uc usecase.%s
	v1.UnimplementedTransportServer
}

var _ v1.TransportServer = (*Transport)(nil)

func NewTransport(uc usecase.%s) *Transport {
	return &Transport{
		uc: uc,
	}
}

%s

`
