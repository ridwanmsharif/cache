

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/ridwanmsharif/cache/practice/protobuf"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement package.GreeterServer.
type server struct{}

// SayHello implements package.GreeterServer
func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetReply, error) {
	return &pb.GreetReply{Message: "GRPC is brilliant! " + in.Name}, nil
}

func (s *server) GreetAgain(ctx context.Context, in *pb.GreetRequest) (*pb.GreetReply, error) {
	return &pb.GreetReply{Message: "Let's move onto the caching service now, shall we? " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
