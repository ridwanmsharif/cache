package main

import (
	"fmt"
	"os"
	"time"

	rpc "github.com/ridwanmsharif/cache/idl"
	"github.com/ridwanmsharif/cache/interceptor"
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
	conn, err := grpc.Dial("localhost:5051",
		grpc.WithTransportCredentials(tlsCreds),
		interceptor.WithClientInterceptor(),
	)

	if err != nil {
		return fmt.Errorf("could not dial %s: %s", "localhost:5051", err)
	}

	cache := rpc.NewCacheClient(conn)

	// Store
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = cache.Store(ctx, &rpc.StoreReq{
		AccountToken: "CLIENT1",
		Key:          "TESTKEY",
		Val:          []byte("TESTVALUE"),
	})

	if err != nil {
		return fmt.Errorf("Failed to store key value pair : %s\n", err)
	}

	// Get
	ctx, _ = context.WithTimeout(context.Background(), 50*time.Millisecond)
	resp, err := cache.Get(ctx, &rpc.GetReq{Key: "TESTKEY"})

	if err != nil {
		return fmt.Errorf("Failed to store key value pair : %s\n", err)
	}

	fmt.Printf("Got value from cache service: %s\n", resp.Val)
	return nil
}

func main() {
	clientMain()
}
