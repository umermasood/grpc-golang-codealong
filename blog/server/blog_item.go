package main

import (
	pb "github.com/umermasood/grpc-golang-codealong/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// we want to create a proto object and the helper function
// this proto object will help us to map the database data inside an object that we can use in our go code
// and the helper function will translate this proto object to a pb.Blog message

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
