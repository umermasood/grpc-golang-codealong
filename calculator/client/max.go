package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating a stream: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending message: %v", req)
			if err = stream.Send(req); err != nil {
				log.Fatalf("Error while sending a req stream: %v", err)
			}

			time.Sleep(1 * time.Second)
		}

		if err = stream.CloseSend(); err != nil {
			log.Fatalf("Error while closing stream: %v\n", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving stream: %v\n", err)
			}

			log.Printf("Received a new max: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
