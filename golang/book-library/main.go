package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ISBN     string `json:"isbn"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ISBN: "123456789", Title: "Tom Sawyer", Author: "Mark Twain", Quantity: 2},
	{ISBN: "097562351", Title: "Old Man and the Sea", Author: "Ernest Hemingway", Quantity: 3},
	{ISBN: "183045751", Title: "The Pearl", Author: "John Steinbeck", Quantity: 4},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	fmt.Println("Welcome to our Library")

	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
