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
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.DELETE("/posts/:id",controllers.PostsDelete)
	router.Run()
}
