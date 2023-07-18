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
	return &pb.HelloReply{Message: "Member " + in.GetName() + " Validated"}, nil
}

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	hostname := os.Getenv("HOSTNAME")
	protocol := os.Getenv("PROTOCOL")

	listener, err := net.Listen(""+protocol+"", ""+hostname+"")
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
