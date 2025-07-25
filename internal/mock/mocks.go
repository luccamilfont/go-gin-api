package mock

import (
	"example/web-service-gin/internal/model"
)

var Client = []model.Client{
	{ID: "1", Name: "Vanessa Clinton"},
	{ID: "2", Name: "Meddler Bill"},
	{ID: "3", Name: "Poltergeist Brown"},
}

var Books = []model.Book{
	{ID: "1", Title: "Blue Train", Author: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Author: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", Price: 39.99},
}
