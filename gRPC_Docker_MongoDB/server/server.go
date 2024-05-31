package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sourabhbiradar/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBlogServiceServer
}

var collection *mongo.Collection // mongo collection
var connStr = "mongodb://admin:pasAdmin@localhost:27017/"
var addr = "0.0.0.0:9321"

func main() {
	//	client, err := mongo.NewClient(options.Client().ApplyURI(connStr)) // deprecated
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogDB").Collection("blog") // looks for blogDB,blog ; if doesn't exist , creates blogDB,blog

	// listen & serve gRPC
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on :%v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v\n", err)
	}

}
