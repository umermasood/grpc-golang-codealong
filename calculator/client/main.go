package main

import (
	pb "github.com/umermasood/grpc-golang-codealong/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err = conn.Close(); err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	doMax(c)
}
