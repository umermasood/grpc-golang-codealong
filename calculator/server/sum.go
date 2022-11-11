package main

import (
	"context"
	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
