package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/mgo.v2/bson"
)

func (s *server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("---deleteteBlog---")
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Can not parse id",
		)
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can not delete object in mongoDB : %v\n", err),
		)
	}

	if res.DeletedCount == 0 { // not deleted since blog not found
		return nil, status.Errorf(
			codes.NotFound,
			"Blog not found",
		)
	}

	return &emptypb.Empty{}, nil
}
