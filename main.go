package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	fmt.Println("Server started....")
	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
