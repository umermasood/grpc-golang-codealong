package main

import (
	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes was invoked with %v", in)

	var k int64 = 2
	var N = in.Number
	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			if err := stream.Send(&pb.PrimeResponse{Result: k}); err != nil {
				log.Fatalf("Couldn't stream response: %v", err)
			}
			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}
	return nil
}
