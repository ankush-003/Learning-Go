package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

// fields are uppercase so that they are exported and accessible outside the package
// `json:"id"` is a struct tag that tells the json package how to serialize the field
type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

// slice of books
var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 5},
	{ID: "3", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
}

// *gin.Context is the context of the current request. It holds all the data (its a pointer to a struct)
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // gives a json response with indentation, http status code 200
}

func createBook(c *gin.Context) {
	var newBook book

	// call BindJSON to bind the received JSON to newBook
	if err := c.BindJSON(&newBook); err != nil {
		// BindJson method uses the Content-Type header to determine how to decode the request body
		// if the Content-Type header is missing or wrong, the method returns an error
		// the error is returned to the client as a 400 Bad Request response
		return
	}

	// add the new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	b, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, b)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id parameter"})
		return
	}
	b, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if b.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not available"})
		return
	}
	b.Quantity--
	c.IndentedJSON(http.StatusOK, b)
}

func returnBook(c *gin.Context) {
	id, ok :=  c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id parameter"})
		return
	}
	b, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	b.Quantity++
	c.IndentedJSON(http.StatusOK, b)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/books/checkout", checkOutBook)
	router.PATCH("/books/return", returnBook)
	router.Run("localhost:8080")
}