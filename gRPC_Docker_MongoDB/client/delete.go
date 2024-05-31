package main

import (
	"context"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting Blog")
	}

	log.Println("Blog Deleted successfully")
}
