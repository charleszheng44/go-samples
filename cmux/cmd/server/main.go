package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "cmuxex/pkg/helloworld"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	httpHandler = http.HandlerFunc(helloHandler)
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello http"))
}

func main() {
	flag.Parse()

	// Create the main listener.
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a cmux.
	m := cmux.New(lis)

	// Match connections in order:
	// First grpc, then HTTP.
	grpcL := m.Match(cmux.HTTP2())
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := grpc.NewServer()
	pb.RegisterGreeterServer(grpcS, &pb.GreetServer{})

	httpS := &http.Server{
		Handler: httpHandler,
	}

	log.Printf("server listening at %v", lis.Addr())
	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	// Start serving!
	m.Serve()
}
