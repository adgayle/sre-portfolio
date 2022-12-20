// book-library provides basic functions to simulate a library of books.
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type Book struct {
	ISBN     string `json:"isbn" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

// books initial library of books.
// TODO: move to document database.
var books = []Book{
	{ISBN: "123456789", Title: "Tom Sawyer", Author: "Mark Twain", Quantity: 2},
	{ISBN: "097562351", Title: "Old Man and the Sea", Author: "Ernest Hemingway", Quantity: 3},
	{ISBN: "183045751", Title: "The Pearl", Author: "John Steinbeck", Quantity: 4},
	{ISBN: "111111111", Title: "The Silver Sword", Author: "Ian Serraillier", Quantity: 0},
}

// GetBooks lists all the books available in the library.
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// CheckoutBook takes out a book from the library.
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

// ReturnBook returns a book to the library.
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

// DeleteBook deletes a book from the library.
func DeleteBook(c *gin.Context) {
	isbn, ok := c.GetQuery("isbn")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing isbn query parameter"})
		return
	}

	bookToDelete, err := bookByISBN(isbn)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	i := slices.Index(books, *bookToDelete)
	books[i] = books[len(books)-1]
	books = books[:len(books)-1]
}

// bookByISBN finds a book given the ISBN.
func bookByISBN(isbn string) (*Book, error) {
	for index, book := range books {
		if book.ISBN == isbn {
			return &books[index], nil
		}
	}

	return nil, fmt.Errorf("book with ISBN %v not found", isbn)
}

// AddBook adds a book to the library.
func AddBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad book JSON"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// GetBookByISBN list a book given the ISBN.
func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	book, err := bookByISBN(isbn)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// HandleHomePage handles GET requests to /.
func HandleHomePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to our Book Library"})
}
