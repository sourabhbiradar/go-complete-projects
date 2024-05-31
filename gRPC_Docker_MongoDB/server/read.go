package main

import (
	"context"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (s *server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Println("---readBlog---")
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"can not parse id",
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with provided ID",
		)
	}

	return newBlogItem(data), nil
}
