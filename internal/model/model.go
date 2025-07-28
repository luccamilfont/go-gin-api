package model

import "time"

type Client struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID        uint       `json:"id"`
	ClientID  int        `json:"client_id"`
	CreatedAt time.Time  `json:"created_at"`
	PaidAt    *time.Time `json:"paid_at"`
	Products  []Product  `gorm:"many2many:order_products;"`
}

type Product struct {
	ID       uint    `json:"id"`
	Barcode  string  `json:"barcode"`
	Price    float32 `json:"price" gorm:"type:decimal(10,2)"`
	ItemID   uint    `json:"item_id"`
	ItemType string  `json:"item_type"`
}

type Book struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	Edition   uint      `json:"edition"`
	Genre     string    `json:"genre"`
	Products  []Product `gorm:"polymorphic:Item;"`
}

type Magazine struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title"`
	Edition  uint      `json:"edition"`
	Company  string    `json:"company"`
	Products []Product `gorm:"polymorphic:Item;"`
}
