package main

// Imports
import (
	"fmt"
	"net"
	"os"

	"github.com/grpc/grpc-go/status"
	rpc "github.com/ridwanmsharif/cache/idl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/reflection"
)

// Server entry point
func accountsMain() {
	if err := runAccountServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run accounts service: %s\n", err)
		os.Exit(1)
	}
}

// Run the server after establishing gRPC connections
func runAccountServer() error {
	tlsCreds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")

	if err != nil {
		return err
	}

	srv := grpc.NewServer(grpc.Creds(tlsCreds))
	as := AccountsService{
		accounts: make(map[string]int64),
	}

	as.accounts["CLIENT1"] = 10
	as.accounts["CLIENT2"] = 10
	as.accounts["CLIENT3"] = 10

	rpc.RegisterAccountsServer(srv, &as)
	l, err := net.Listen("tcp", "localhost:5052")
	if err != nil {
		return err
	}

	// Blocks until complete
	return srv.Serve(l)
}

type AccountsService struct {
	accounts map[string]int64
}

// Get Method
func (s *AccountsService) GetByToken(ctx context.Context, req *rpc.GetByTokenReq) (*rpc.GetByTokenResp, error) {
	val, ok := s.accounts[req.Token]
	if !ok {
		// err := "Key not found :%s\n"
		// return nil, fmt.Errorf(err, req.Key)
		return nil, status.Errorf(codes.NotFound, "Key not found: %s\n", req.Token)
	}
	return &rpc.GetByTokenResp{Allowed: val}, nil
}

func main() {
	accountsMain()
}
