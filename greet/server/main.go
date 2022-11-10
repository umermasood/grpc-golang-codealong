package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/umermasood/grpc-golang-codealong/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
