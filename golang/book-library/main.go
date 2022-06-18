// book-library provides basic functions to simulate a library of books
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ISBN     string `json:"isbn"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// books initial library of books
// TODO: move to document database
var books = []Book{
	{ISBN: "123456789", Title: "Tom Sawyer", Author: "Mark Twain", Quantity: 2},
	{ISBN: "097562351", Title: "Old Man and the Sea", Author: "Ernest Hemingway", Quantity: 3},
	{ISBN: "183045751", Title: "The Pearl", Author: "John Steinbeck", Quantity: 4},
}

// GetBooks lists all the books available in the library
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// CheckoutBook takes out a book from the library
func CheckoutBook(c *gin.Context) {
	isbn, ok := c.GetQuery("isbn")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing isbn query parameter"})
		return
	}

	book, err := bookByISBN(isbn)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// ReturnBook returns a book to the library
func ReturnBook(c *gin.Context) {
	isbn, ok := c.GetQuery("isbn")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing isbn query parameter"})
		return
	}

	book, err := bookByISBN(isbn)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, gin.H{"message": book})
}

// bookByISBN finds a book given the ISBN
func bookByISBN(isbn string) (*Book, error) {
	for index, book := range books {
		if book.ISBN == isbn {
			return &books[index], nil
		}
	}

	return nil, fmt.Errorf("book with ISBN %v not found", isbn)
}

// AddBook adds a book to the library
func AddBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// GetBookByISBN list a book given the ISBN
func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	book, err := bookByISBN(isbn)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// handleHomePage handles GET requests to /
func HandleHomePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to our Book Library"})
}

// main the controlling function
func main() {

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", HandleHomePage)
	router.GET("/books", GetBooks)
	router.POST("/books", AddBook)
	router.GET("/books/:isbn", GetBookByISBN)
	router.PATCH("/checkout", CheckoutBook)
	router.PATCH("/return", ReturnBook)
	router.Run("localhost:8080")
}
