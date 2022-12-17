package main

import (
	"context"
	"log"
	"net"

	apiv1 "github.com/charconstpointer/ihateannotations/proto/gen/go/api/v1"
	"google.golang.org/grpc"
)

func main() {
	s := Server{}
	server := grpc.NewServer([]grpc.ServerOption{}...)
	apiv1.RegisterApiServiceServer(server, s)
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := server.Serve(ln); err != nil {
		log.Fatal(err.Error())
	}
}

type Server struct {
}

func (s Server) Foo(ctx context.Context, request *apiv1.FooRequest) (*apiv1.FooResponse, error) {
	return &apiv1.FooResponse{}, nil
}
