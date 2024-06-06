package main

import (
	"find-a-friend/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.ConnectDb()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.CreateTables()
	if err != nil {
		log.Fatalf(err.Error())
	}

	router := gin.Default()
	// initialize routes
	log.Println("🚀 Server starting...	")
	if err = router.Run(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
