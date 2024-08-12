package main

import (
	"abhiroopsanta.dev/event-booking-api/db"
	"abhiroopsanta.dev/event-booking-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDb()
	server := gin.Default()
	fmt.Println("Server started....")

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err = server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request object"})
		return
	}

	//event.Id = 1
	//event.UserId = 1
	err = event.Save()
	if err != nil {
		fmt.Printf("Error saving event: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "data": event})
}
