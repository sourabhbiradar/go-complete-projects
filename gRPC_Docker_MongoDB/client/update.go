package main

import (
	"context"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Maxi",
		Title:    "First training",
		Content:  "Updating & adding to previous Blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error while updating :%v\n", err)
	}

	log.Println("Blog Updated")
}
