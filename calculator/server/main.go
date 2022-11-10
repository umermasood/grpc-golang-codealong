package main

import (
	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.SumServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
