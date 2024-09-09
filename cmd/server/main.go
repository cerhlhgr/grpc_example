package main

import (
	"context"
	desc "github.com/cerhlhgr/grpc_example/pkg/pack_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	desc.UnimplementedPackServiceServer
}

func (s *server) Get(ctx context.Context, pack *desc.PackRequest) (*desc.PackResponse, error) {
	return &desc.PackResponse{Name: "Имя", Description: "Описал"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterPackServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
