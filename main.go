package main

import (
	"find-a-friend/db"
	"find-a-friend/http/controllers"
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
	controllers.NewOrgController(router)
	controllers.NewPetController(router)

	log.Println("🚀 Server starting...	")
	if err = router.Run(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
