package main

import (
	"log"

	"go-gin-api/internal/api"
	"go-gin-api/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsnURI := "mydb.sqlite"
	db, err := gorm.Open(sqlite.Open(dsnURI), &gorm.Config{})
	db.AutoMigrate(&model.Client{}, &model.Order{}, &model.Product{}, &model.Book{}, &model.Magazine{})
	if err != nil {
		log.Fatal(err)
	}

	router := api.FnRouter(db)
	router.Run("localhost:8080")
}
