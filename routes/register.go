package routes

import (
	"abhiroopsanta.dev/event-booking-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id, expecting an integer"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		log.Printf("Error getting event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event, please try again later"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		log.Printf("Error registering for event with id %v: %v\n", eventId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event, please try again later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registered for event successfully", "data": event})
}

func cancelRegistration(context *gin.Context) {

}

func getRegisteredEvents(context *gin.Context) {

}
