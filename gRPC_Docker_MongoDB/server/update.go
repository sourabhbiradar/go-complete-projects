package main

import (
	"context"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/mgo.v2/bson"
)

func (s *server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Println("---updateBlog---")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cound not parse ID",
		)
	}

	data := &BlogItem{
		Author_Id: in.AuthorId,
		Title:     in.Title,
		Content:   in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data}, // all fileds == $set
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"could not update",
		)
	}

	if res.MatchedCount == 0 { // no entries with given id
		return nil, status.Errorf(
			codes.NotFound,
			"could not find blog with id",
		)
	}

	return &emptypb.Empty{}, nil
}
