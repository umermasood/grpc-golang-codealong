package main

import (
	"context"
	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
	"log"
)

func doSum(c pb.SumServiceClient) {
	log.Printf("doSum was invoked\n")
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Couldn't perfrom Sum: %v\n", err)
	}
	log.Printf("Sum: %v\n", resp.Result)
}
