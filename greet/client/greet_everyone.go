package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/umermasood/grpc-golang-codealong/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Umer"},
		{FirstName: "Masood"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			if err = stream.Send(req); err != nil {
				log.Fatalf("Error while sending stream: %v\n", err)
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
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
