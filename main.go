package main

import (
	"errors"
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

/* Helper functions */

func getBookById(id string) (*book, error) {
	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("no book is found")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	// create a var of type book
	// will be initialized with falsy values str => "", int => 0
	var newBook book

	// write the json body to the variable newBook
	// this will not validate or check type of body extra or missing keys will not panic
	// But BindJson will panic if the key exist with the wrong type

	// If a field exists in the JSON Body, it will get parsed then update the newBook var with the value sent in the Json body
	// Otherwise all other fields will not be changes
	// extra fields will be discarded

	// Note that you pass the memory address of the var so it can be updated by the BindJSON function
	// if the type is wrong or any other error occur the err object will not be nil
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// append a var to slice books
	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBook(c *gin.Context) {
	// Get params from gin context
	book, err := getBookById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/book", addBook)
	router.GET("/book/:id", getBook)

	// Start Server
	router.Run("localhost:8080")
}
