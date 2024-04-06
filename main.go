package main

import (
	"log"
	"postgres-gorm/controllers"
	"postgres-gorm/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)
	router.Run()
}
