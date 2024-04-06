package main

import (
	"log"
	"postgres-gorm/database"
	"postgres-gorm/models"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Post{})
}
