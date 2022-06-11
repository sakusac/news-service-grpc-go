package models

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"news-service-grpc-go/config"
)

var Db *sql.DB
var err error

func connectDB() (db *sql.DB, err error) {
	sqlDriver := fmt.Sprintf("%v", config.Config.SQLDriver)
	connInfo := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		config.Config.DBHost,
		config.Config.DbPort,
		config.Config.DbUser,
		config.Config.DbName,
		config.Config.DbPassword,
	)

	Db, err = sql.Open(sqlDriver, connInfo)

	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	return Db, err
}
