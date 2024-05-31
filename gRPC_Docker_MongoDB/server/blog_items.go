package main

import (
	pb "github.com/sourabhbiradar/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The BlogItem struct defines the fields for a blog item, matching how it would be stored in MongoDB.
// BlogItem represents a blog item stored in MongoDB.
type BlogItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Author_Id string             `bson:"author_id"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
}

// converts the BlogItem fields into the corresponding fields of the Protocol Buffers Blog message.
func newBlogItem(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.Author_Id,
		Title:    data.Title,
		Content:  data.Content,
	}
}
