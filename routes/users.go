package routes

import (
	"abhiroopsanta.dev/event-booking-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request object"})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Printf("Error in user signup: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to sign up user, please try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "data": user})
}
