package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"news-service-grpc-go/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNewsServiceClient(conn)
	callListNews(client)
}

func callListNews(client pb.NewsServiceClient) {
	res, err := client.ListNews(context.Background(), &pb.NewsRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetNews())
}
