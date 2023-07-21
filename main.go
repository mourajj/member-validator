package main

import (
	"context"
	"fmt"
	"log"
	"membervalidator/pb"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Member " + in.GetName() + " being processed and validated")
	//TODO: validation logic
	return &pb.HelloReply{Message: "Member " + in.GetName() + " Validated"}, nil
}

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	protocol := os.Getenv("PROTOCOL")

	listener, err := net.Listen(""+protocol+"", ":9000")
	if err != nil {
		panic(err)
	}
	println("Running member validator gRPC Server")

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve", err)
	}

}
