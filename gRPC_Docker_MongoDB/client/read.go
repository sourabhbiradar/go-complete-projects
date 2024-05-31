package main

import (
	"context"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog---")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading : %v\n", err)
	}

	log.Printf("Blog :%v\n", res)
	return res
}
