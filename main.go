package main

import (
	"context"
	v1 "github.com/charconstpointer/ihateannotations/proto/gen/go/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	c := v1.NewApiServiceClient(conn)
	res, err := c.Foo(context.Background(), &v1.FooRequest{})
	if err != nil {
		log.Fatalf("Failed to call Foo: %v", err)
	}
	log.Println(res)
}
