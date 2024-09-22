package middlewares

import (
	"abhiroopsanta.dev/event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing authorization token"})
		return
	}

	authToken := strings.Split(authHeader, "Bearer ")[1]

	userId, err := utils.VerifyToken(authToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
