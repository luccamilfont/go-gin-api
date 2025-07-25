package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example/web-service-gin/internal/mock"
	"example/web-service-gin/internal/model"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mock.Books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range mock.Books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func PostBooks(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	mock.Books = append(mock.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
