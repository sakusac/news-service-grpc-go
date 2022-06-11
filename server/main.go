package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"news-service-grpc-go/models"
	"news-service-grpc-go/pb"
)

type server struct {
	pb.UnimplementedNewsServiceServer
}

func (*server) ListNews(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	fmt.Println("ListNews was invoked.")

	// get data from db
	news, _ := models.GetNews()

	res := &pb.NewsResponse{
		News: news,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNewsServiceServer(s, &server{})
	fmt.Println("server is running...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
