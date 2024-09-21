package utils

import (
	"errors"
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
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET environment variable not set")
	}
	log.Println("JWT Secret retrieved: ", jwtSecret)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			return nil, errors.New("JWT_SECRET environment variable not set")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("could not parse claims")
	}

	expiryTime := time.Unix(int64(claims["exp"].(float64)), 0)

	if time.Now().After(expiryTime) {
		return errors.New("token has expired")
	}

	return nil
}
