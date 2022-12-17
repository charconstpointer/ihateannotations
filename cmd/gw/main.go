package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"

	apiv1 "github.com/charconstpointer/ihateannotations/proto/gen/go/api/v1"
	v1 "github.com/charconstpointer/ihateannotations/proto/gen/openapiv2/api/v1"
	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

// can handle both grpc and rest on the same port
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	flag.Parse()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gmux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", gmux)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(v1.SwaggerSpec)))

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := apiv1.RegisterApiServiceHandlerFromEndpoint(ctx, gmux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}
	http.ListenAndServe(":8081", mux)
}
