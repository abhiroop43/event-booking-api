package routes

import (
	"abhiroopsanta.dev/event-booking-api/models"
	"abhiroopsanta.dev/event-booking-api/utils"
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

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request object"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate token, please try again later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token, "message": "login successful"})
}
