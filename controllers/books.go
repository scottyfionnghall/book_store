package controllers

import (
	"net/http"

	"gintest/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ReleaseYear string `json:"release_year" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
	AuthorID    uint
}

func FindBooks(c *gin.Context) {
	var books []models.Book

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit key"})
		return
	}

	order := strings.ToUpper(c.DefaultQuery("order", "asc"))
	if order != "ASC" && order != "DESC" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order key"})
		return
	}

	by := c.DefaultQuery("by", "id")
	if by != "id" && by != "title" && by != "release_year" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid by key"})
		return
	}

	models.DB.Limit(limit).Order(by + " " + order).Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	var author models.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("name", input.Author).First(&author).Error; err != nil {
		author = models.Author{Name: input.Author}
		models.DB.Create(&author)
	}
	book := models.Book{Title: input.Title, ReleaseYear: input.ReleaseYear, AuthorID: author.ID}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	var author models.Author

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("name", input.Author).First(&author).Error; err != nil {
		author = models.Author{Name: input.Author}
		models.DB.Create(&author)
	}
	input.AuthorID = author.ID
	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindAuthors(c *gin.Context) {
	var author []models.Author

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit key"})
		return
	}

	order := strings.ToUpper(c.DefaultQuery("order", "asc"))
	if order != "ASC" && order != "DESC" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order key"})
		return
	}

	by := c.DefaultQuery("by", "id")
	if by != "first_name" && by != "last_name" && by != "id" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid by key"})
		return
	}

	models.DB.Limit(limit).Order(by + " " + order).Find(&author)

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func FindAuthor(c *gin.Context) {
	var author models.Author

	if err := models.DB.Where("id=?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func FindAuthorWorks(c *gin.Context) {
	var author models.Author
	var books []models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := models.DB.Where("author_id=?", author.ID).Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not find works of that author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})

}
