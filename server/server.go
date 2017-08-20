package main

// Imports
import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/grpc/grpc-go/status"
	rpc "github.com/ridwanmsharif/cache/idl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/reflection"
)

// Server entry point
func serverMain() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run caching service: %s\n", err)
		os.Exit(1)
	}
}

// Run the server after establishing gRPC connections
func runServer() error {
	tlsCreds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")

	if err != nil {
		return err
	}

	srv := grpc.NewServer(grpc.Creds(tlsCreds))

	// Create the client TLS credentials
	tlsCreds2, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
	if err != nil {
		return fmt.Errorf("could not load tls cert: %s", err)
	}

	// Establish a connection with Accounts server
	conn, err := grpc.Dial("localhost:5052", grpc.WithTransportCredentials(tlsCreds2))
	if err != nil {
		return fmt.Errorf("could not dial %s: %s", "localhost:5052", err)
	}

	account := rpc.NewAccountsClient(conn)

	cs := CacheService{
		accounts:      account,
		store:         make(map[string][]byte),
		keysByAccount: make(map[string]int64),
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
	accounts      rpc.AccountsClient
	store         map[string][]byte
	keysByAccount map[string]int64
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
	// Logging
	start := time.Now()

	resp, err := s.accounts.GetByToken(context.Background(), &rpc.GetByTokenReq{
		Token: req.AccountToken,
	})
	log.Printf("Accounts GetByToken duration: %s\n", time.Since(start))

	if err != nil {
		return nil, err
	}

	if s.keysByAccount[req.AccountToken] >= resp.Allowed {
		return nil, status.Errorf(codes.FailedPrecondition, "Account %s exceeded max key limit: %d\n", req.AccountToken, resp.Allowed)
	}

	if _, ok := s.store[req.Key]; !ok {
		s.keysByAccount[req.AccountToken] += 1
	}

	s.store[req.Key] = req.Val
	return &rpc.StoreResp{}, nil
}

func main() {
	serverMain()
}
