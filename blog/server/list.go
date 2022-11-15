package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/umermasood/grpc-golang-codealong/blog/proto"
)

func (*Server) ListBlogs(_ *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	ctx := context.Background()
	cur, err := collection.Find(ctx, primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatalf("Error while closing cursor: %v\n", err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		err = stream.Send(documentToBlog(data))
		if err != nil {
			log.Fatalf("Error while streaming data: %v", err)
		}
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}
