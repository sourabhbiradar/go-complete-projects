package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sourabhbiradar/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("---listBlog---")
	strm, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs : %v\n", err)
	}

	for {
		res, err := strm.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened : %v\n", err)
		}
		log.Println(res)
	}
}
