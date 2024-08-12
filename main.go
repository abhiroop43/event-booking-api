package main

import (
	"abhiroopsanta.dev/event-booking-api/db"
	"abhiroopsanta.dev/event-booking-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDb()
	server := gin.Default()
	fmt.Println("Server started....")

	routes.RegisterRoutes(server)

	err = server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
