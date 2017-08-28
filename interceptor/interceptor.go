package interceptor

import (
	"log"
	"time"

	"github.com/grpc/grpc-go/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	var (
		start    = time.Now()
		attempts = 0
		err      error
		backTime time.Duration = 10
		backoff                = time.NewTimer(time.Millisecond * backTime)
	)

	for {
		attempts += 1
		select {
		case <-ctx.Done():
			err = status.Errorf(codes.DeadlineExceeded, "timeout reached before retry attempt")
		case <-backoff.C:
			log.Println("Inside For loop")
			err = invoker(ctx, method, req, reply, cc, opts...)
			backTime *= 2

			if err != nil {
				log.Printf("Invoke method=%s duration=%s error=%v", method, time.Since(start), err)

				if !backoff.Stop() {
					<-backoff.C
				}
				backoff.Reset(time.Millisecond * backTime)
				continue
			}
		}
		break
	}
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
