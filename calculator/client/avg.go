package main

import (
	"context"
	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg function was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doAvg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("doAvg: %v\n", res.Result)
}
