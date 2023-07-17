package main

import (
	"context"
	"log"
	"membervalidator/pb"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Member " + in.GetName() + " Validated"}, nil
}

func main() {
	println("Running member validator gRPC Server")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve", err)
	}
}
