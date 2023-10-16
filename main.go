package main

import (
	"log"
	"net/http"
	"os"

	"gintest/controllers"
	"gintest/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()

	err := godotenv.Load(".env")

	err = models.ConnectDatabase(os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
		return
	}

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.GET("/author", controllers.FindAuthors)
	r.GET("/author/:id", controllers.FindAuthor)
	r.GET("/books/by_author/:id", controllers.FindAuthorWorks)
	r.Run()
}
