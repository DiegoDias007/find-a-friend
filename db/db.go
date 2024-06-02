package db

import (
	"context"
	"find-a-friend/utils"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn
var err error

func ConnectDb() error {
	utils.LoadEnv()
	databaseURL := os.Getenv("DATABASE_URL")
	DB, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return err
	}
	fmt.Println("Connected to the database.")
	return nil
}

func CloseDb() {
	if DB != nil {
		DB.Close(context.Background())
	}
}
