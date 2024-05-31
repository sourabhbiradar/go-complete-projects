package main

import (
	"context"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog---")

	// create blog
	blog := &pb.Blog{
		AuthorId: "Green",
		Title:    "First Practice With RCB",
		Content:  "Greet experience , got to know teammates",
	}

	// send to server
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error : %v\n", err)
	}

	log.Printf("Blog has been created :%s\n", res.Id)
	return res.Id
}
