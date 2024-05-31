package main

import (
	"log"

	pb "github.com/sourabhbiradar/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:9321"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// createBlog(c)   // create blog

	//id := createBlog(c)   // read blog
	//readBlog(c, id)       // with vaild id
	//readBlog(c, "387538") // with invalid id

	//updateBlog(c, id) // update blog

	listBlogs(c) // list all blogs

	//deleteBlog(c, id)
}
