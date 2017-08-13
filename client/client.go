package main

import (
	"fmt"
	// "net"
	"os"

	rpc "github.com/ridwanmsharif/cache/idl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/reflection"
)

// Client entry point
func clientMain() {
	if err := runClient(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %s\n", err)
		os.Exit(1)
	}
}

// Run Client
func runClient() error {

	// Create the client TLS credentials
	tlsCreds, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
	if err != nil {
		return fmt.Errorf("could not load tls cert: %s", err)
	}

	// Create a connection with the TLS credentials
	conn, err := grpc.Dial("localhost:5051", grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		return fmt.Errorf("could not dial %s: %s", "localhost:5051", err)
	}

	cache := rpc.NewCacheClient(conn)

	// Store
	_, err = cache.Store(context.Background(), &rpc.StoreReq{Key: "TESTKEY", Val: []byte("TESTVALUE")})

	if err != nil {
		return fmt.Errorf("Failed to store key value pair : %s\n", err)
	}

	// Get
	resp, err := cache.Get(context.Background(), &rpc.GetReq{Key: "TESTKEY"})

	if err != nil {
		return fmt.Errorf("Failed to store key value pair : %s\n", err)
	}

	fmt.Printf("Got value from cache service: %s\n", resp.Val)
	return nil
}

func main() {
	clientMain()
}
