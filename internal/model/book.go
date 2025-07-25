package model

type Client struct {
	ID   string `json:"id"`
	Name string `json:"title"`
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
