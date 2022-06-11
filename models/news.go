package models

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"news-service-grpc-go/pb"
	"time"
)

func GetNews() (News []*pb.News, err error) {
	fmt.Println("get news was invoked")
	ctx := context.Background()

	// Get News data from DB
	Db, err = connectDB()
	defer Db.Close()

	rows, err := Db.QueryContext(ctx, `SELECT id, title, content, created_at FROM news;`)
	if err != nil {
		log.Fatalf("get data error %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id             int32
			title, content string
			createdAt      time.Time
		)

		if err := rows.Scan(&id, &title, &content, &createdAt); err != nil {
			log.Fatalf("scan the news: %v", err)
		}

		convCreatedAt := timestamppb.New(createdAt)
		News = append(News, &pb.News{
			Id:        id,
			Title:     title,
			Content:   content,
			CreatedAt: convCreatedAt,
		})
	}

	rows.Close()
	return News, err
}
