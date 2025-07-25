package api

import (
	"github.com/gin-gonic/gin"
)

func FnRouter() {
	router := gin.Default()
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBookByID)
	router.POST("/books", PostBooks)

	router.Run("localhost:8080")
}
