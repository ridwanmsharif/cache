package main

// Imports
import (
	"fmt"
	"net"
	"os"

	// "github.com/pkg/errors"
	"github.com/grpc/grpc-go/status"
	rpc "github.com/ridwanmsharif/cache/idl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/reflection"
)

// Server entry point
func serverMain() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run caching service: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Server running. Caching service started.")
}

// Run the server after establishing gRPC connections
func runServer() error {
	srv := grpc.NewServer()
	cs := CacheService{
		store: make(map[string][]byte),
	}
	rpc.RegisterCacheServer(srv, &cs)
	l, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		return err
	}

	// Blocks until complete
	return srv.Serve(l)
}

type CacheService struct {
	store map[string][]byte
}

// Get Method
func (s *CacheService) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	val, ok := s.store[req.Key]
	if !ok {
		// err := "Key not found :%s\n"
		// return nil, fmt.Errorf(err, req.Key)
		return nil, status.Errorf(codes.NotFound, "Key not found: %s\n", req.Key)
	}
	return &rpc.GetResp{Val: val}, nil
}

// Store Method
func (s *CacheService) Store(ctx context.Context, req *rpc.StoreReq) (*rpc.StoreResp, error) {
	s.store[req.Key] = req.Val
	return &rpc.StoreResp{}, nil
}

func main() {
	serverMain()
}
