package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"time"

	pb "cmuxex/pkg/helloworld"

	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet again: %v", err)
	}
	log.Printf("Greeting Again: %s", r.GetMessage())

	rep, err := http.Get("http://localhost:50051")
	if err != nil {
		log.Fatalf("could not say hello through http: %s", err)
	}
	repStr, err := io.ReadAll(rep.Body)
	if err != nil {
		log.Fatalf("fail to read the response body: %s", err)
	}

	log.Printf("Hello http: %s", repStr)
}
