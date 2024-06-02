package main

import (
	"find-a-friend/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.ConnectDb()
	if err != nil {
		log.Fatal("Error when connecting to the database.")
	}
	router := gin.Default()
	// initialize routes
	err = router.Run()
	if err != nil {
		log.Fatalf("Error starting server.")
	}
}
