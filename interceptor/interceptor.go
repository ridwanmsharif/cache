package interceptor

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Wrap the client interceptor in a Unary interceptor
func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

// Define the client interceptor
func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Invoke method=%s duration=%s error=%v", method, time.Since(start), err)
	return err
}

// Wrap the server interceptor in a Unary interceptor
func ServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// Define the server interceptor
func serverInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("Invoke method=%s duration=%s error=%v", info.FullMethod, time.Since(start), err)
	return resp, err
}
