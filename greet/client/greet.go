package main

import (
	"context"
	pb "github.com/umermasood/grpc-golang-codealong/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Umer",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
