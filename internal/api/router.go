package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FnRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		GetProducts(c, db)
	})
	router.GET("/books", func(c *gin.Context) {
		GetBooks(c, db)
	})
	router.GET("/books/:id", func(c *gin.Context) {
		GetBookByID(c, db)
	})
	router.POST("/newBook", func(c *gin.Context) {
		PostBook(c, db)
	})
	router.POST("/newProduct", func(c *gin.Context) {
		PostProduct(c, db)
	})

	router.Run("localhost:8080")
	return router
}
