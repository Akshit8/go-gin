package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	server.Run(port)
}