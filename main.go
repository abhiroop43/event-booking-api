package main

import (
	"abhiroopsanta.dev/event-booking-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	fmt.Println("Server started....")

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
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

	event.Id = 1
	event.UserId = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "data": event})
}
