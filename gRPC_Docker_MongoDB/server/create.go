package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Println("---createBlog---")
	// create blogItem instance >> insert in DB
	data := BlogItem{
		Author_Id: in.AuthorId,
		Title:     in.Title,
		Content:   in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v\n", err),
		)
	}

	// convert returned ID to primitive.ObjID & check if it was success
	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"could not convert to oid",
		)
	}

	// if everything is ok return oid to client
	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
