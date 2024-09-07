package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{Id: "1", Title: "One Title", Author: "One", Quantity: 2},
	{Id: "2", Title: "Two Title", Author: "Two", Quantity: 55},
	{Id: "3", Title: "Three Title", Author: "Three", Quantity: 5},
	{Id: "4", Title: "Four Title", Author: "Four", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
