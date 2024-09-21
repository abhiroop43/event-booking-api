package routes

import (
	"abhiroopsanta.dev/event-booking-api/models"
	"abhiroopsanta.dev/event-booking-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get events, please try again later"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "events retrieved successfully", "data": events})
}

func createEvent(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	log.Println("Auth header: ", authHeader)

	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "missing authorization token"})
		return
	}

	authToken := strings.Split(authHeader, "Bearer ")[1]
	log.Println("Auth token: ", authToken)

	err := utils.VerifyToken(authToken)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var event models.Event

	err = context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request object"})
		return
	}

	err = event.Save()
	if err != nil {
		fmt.Printf("Error saving event: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event, please try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "data": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id, expecting an integer"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Printf("Error getting event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event, please try again later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event retrieved successfully", "data": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id, expecting an integer"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		fmt.Printf("Error getting event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event, please try again later"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request object"})
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.Update()
	if err != nil {
		fmt.Printf("Error getting event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event, please try again later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event created successfully", "data": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id, expecting an integer"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		fmt.Printf("Error getting event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event, please try again later"})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		fmt.Printf("Error deleting event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event, please try again later"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
