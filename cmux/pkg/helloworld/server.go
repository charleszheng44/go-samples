package helloworld

import (
	"context"
	"log"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (s *GreetServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *GreetServer) SayHelloAgain(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello Again " + in.GetName()}, nil
}
