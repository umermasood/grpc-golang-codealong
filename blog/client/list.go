package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/umermasood/grpc-golang-codealong/blog/proto"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
