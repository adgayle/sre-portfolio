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

func checkoutBook(context *gin.Context) {
	isbn, ok := context.GetQuery("isbn")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing isbn query parameter"})
		return
	}

	book, err := bookByISBN(isbn)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}
	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, book)
}

func returnBook(context *gin.Context) {
	isbn, ok := context.GetQuery("isbn")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing isbn query parameter"})
		return
	}

	book, err := bookByISBN(isbn)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	context.IndentedJSON(http.StatusOK, gin.H{"message": book})
}

func bookByISBN(isbn string) (*book, error) {
	for index, book := range books {
		if book.ISBN == isbn {
			return &books[index], nil
		}
	}

	return nil, fmt.Errorf("book with ISBN %v not found", isbn)
}

func addBook(context *gin.Context) {
	var newBook book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByISBN(context *gin.Context) {
	isbn := context.Param("isbn")
	book, err := bookByISBN(isbn)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	fmt.Println("Welcome to our Book Library")

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:isbn", getBookByISBN)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
