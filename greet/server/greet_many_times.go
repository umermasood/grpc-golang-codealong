package main

import (
	"fmt"
	"log"

	pb "github.com/umermasood/grpc-golang-codealong/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		if err := stream.Send(&pb.GreetResponse{Result: res}); err != nil {
			log.Fatalf("Couldn't stream response: %v", err)
		}
	}

	return nil
}
