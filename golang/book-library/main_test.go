package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	return router
}

// prettyJSON formats the JSON string using indent
func prettyJSON(str string, indent string) (string, error) {
	var pretty bytes.Buffer

	if err := json.Indent(&pretty, []byte(str), "", indent); err != nil {
		return "", err
	}

	return pretty.String(), nil
}

func TestHandleHomePage(t *testing.T) {
	mockResponse := `{"message": "Welcome to our Book Library"}`
	res, err := prettyJSON(mockResponse, "    ")
	if err != nil {
		log.Fatal(err)
	}
	r := SetUpRouter()
	r.GET("/", HandleHomePage)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, res, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAddBook(t *testing.T) {
	r := SetUpRouter()
	r.POST("/books", AddBook)
	book := Book{
		ISBN:     "097562351",
		Title:    "Old Man and the Sea",
		Author:   "Ernest Hemingway",
		Quantity: 3,
	}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBooks(t *testing.T) {
	r := SetUpRouter()
	r.GET("/books", GetBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var books []Book
	json.Unmarshal(w.Body.Bytes(), &books)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
}

func TestCheckoutBook(t *testing.T) {
	r := SetUpRouter()

	// Book will be found
	r.PATCH("/books", CheckoutBook)
	reqFound, _ := http.NewRequest("PATCH", "/books?isbn=123456789", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Equal(t, http.StatusOK, w.Code)

	// Book will not be found
	reqNotFound, _ := http.NewRequest("PATCH", "/books?isbn=000000000", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Bad request missing isbn
	reqBadNoISBN, _ := http.NewRequest("PATCH", "/books", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqBadNoISBN)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Bad request book not available
	reqBadNotAvailable, _ := http.NewRequest("PATCH", "/books?isbn=111111111", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqBadNotAvailable)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReturnBook(t *testing.T) {
	r := SetUpRouter()
	r.PATCH("/books", ReturnBook)
	reqFound, _ := http.NewRequest("PATCH", "/books?isbn=123456789", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PATCH", "/books?isbn=000000000", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetBookByISBN(t *testing.T) {
	r := SetUpRouter()
	r.GET("/books/:isbn", GetBookByISBN)
	req, _ := http.NewRequest("GET", "/books/123456789", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBookByISBN(t *testing.T) {
	bookFound, _ := bookByISBN("123456789")
	assert.Equal(t, bookFound.ISBN, "123456789")

	isbn := "0000000000"
	_, err := bookByISBN(isbn)
	assert.Error(t, err, "book with ISBN "+isbn+" not found")
}
