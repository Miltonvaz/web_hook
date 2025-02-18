package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.POST("/webhook", func (ctx  *gin.Context){})

	port := os.Getenv("PORT" )

	if port == ""{
		port = "8080"
	}

	router.Run()
}