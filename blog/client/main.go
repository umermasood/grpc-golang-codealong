package main

import (
	pb "github.com/umermasood/grpc-golang-codealong/blog/proto"
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

	c := pb.NewBlogServiceClient(conn)

	//id := createBlog(c)
	//readBlog(c, id) // test readBlog with a valid id
	//readBlog(c, "NonExistingID") // test readBlog with an invalid id
	//updateBlog(c, id)
	listBlog(c)
}
