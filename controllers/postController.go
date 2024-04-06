package controllers

import (
	"postgres-gorm/database"
	"postgres-gorm/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.Status(400)
		return
	}
	post := models.Post{Title: body.Title, Body: body.Body}
	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	database.DB.Find(&posts)

	// Respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get ID of URL
	id := c.Param("id")
	// Get the posts
	var post models.Post
	database.DB.First(&post, id)

	// Respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get ID of URL
	id := c.Param("id")
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.Status(400)
		return
	}
	var post models.Post
	database.DB.First(&post, id)

	database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})

}
