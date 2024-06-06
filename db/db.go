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
		return fmt.Errorf("error when connecting to the database: %v", err)
	}
	fmt.Println("Connected to the database.")
	return nil
}

func CreateTables() error {
	createOrgTable := `
	CREATE TABLE IF NOT EXISTS org (
		id SERIAL PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		address TEXT NOT NULL,
		whatsapp TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err = DB.Exec(context.Background(), createOrgTable)
	if err != nil {
		return fmt.Errorf("error when creating org table: %v", err.Error())
	}

	createPetTable := `
		CREATE TABLE IF NOT EXISTS pet (
			id SERIAL PRIMARY KEY NOT NULL,
			name TEXT NOT NULL,
			city TEXT NOT NULL,
			species TEXT NOT NULL,
			breed TEXT NOT NULL,
			height NUMERIC(3, 2) NOT NULL,
			weight NUMERIC(3, 2) NOT NULL,
			org_id INT REFERENCES org(id),
			UNIQUE(org_id)
		)
	`

	_, err := DB.Exec(context.Background(), createPetTable)
	if err != nil {
		return fmt.Errorf("error when creating pet table: %v", err.Error())
	}

	return nil
}

func CloseDb() {
	if DB != nil {
		DB.Close(context.Background())
	}
}
