package template

var Interceptor = `package transport

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *Transport) ErrInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	statusCode, reason := codes.OK, ""
	reply, err := handler(ctx, req)
	if err != nil {
		log.Printf("[ERROR] %v: %v", info.FullMethod, err)
		switch {
		default:
		}
	}
	return reply, status.Error(statusCode, reason)
}

func (t *Transport) RecoveryInterceptor(p interface{}) (err error) {
	log.Printf("panic occurs becauses of %v", p)
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}
`
