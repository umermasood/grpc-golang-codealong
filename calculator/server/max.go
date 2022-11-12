package main

import (
	"io"
	"log"

	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var max int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving client stream: %v\n", err)
		}

		if req.Number > max {
			max = req.Number

			if err = stream.Send(&pb.MaxResponse{
				Result: max,
			}); err != nil {
				log.Fatalf("Error while sending response: %v\n", err)
			}
		}
	}
}
