package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"example/web-service-gin/internal/model"
)

func GetBooks(c *gin.Context, db *gorm.DB) {
	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var book []model.Book
	if err := db.Find(&book, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func PostBooks(c *gin.Context, db *gorm.DB) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}
