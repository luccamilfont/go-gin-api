package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Client struct {
	Model
	Name string `json:"name"`
}

type Order struct {
	Model
	ClientID int        `json:"client_id"`
	PaidAt   *time.Time `json:"paid_at"`
	Products []OrderProduct
}

type Product struct {
	Model
	Barcode  string  `json:"barcode"`
	Price    float32 `json:"price" gorm:"type:decimal(10,2)"`
	ItemID   uint    `json:"item_id"`
	ItemType string  `json:"item_type"`
	Orders   []OrderProduct
}

type OrderProduct struct {
	Model
	OrderID   uint `gorm:"foreignKey:OrderID;references:ID"`
	ProductID uint `gorm:"foreignKey:ProductID;references:ID"`
	Quantity  uint `json:"quantity"`
}
type Book struct {
	Model
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	Edition   uint      `json:"edition"`
	Genre     string    `json:"genre"`
	Products  []Product `gorm:"polymorphic:Item;"`
}

type Magazine struct {
	Model
	Title    string    `json:"title"`
	Edition  uint      `json:"edition"`
	Company  string    `json:"company"`
	Products []Product `gorm:"polymorphic:Item;"`
}
