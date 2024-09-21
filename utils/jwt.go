package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Println("JWT Secret retrieved: ", jwtSecret)
	return token.SignedString([]byte(os.Getenv(jwtSecret)))
}
