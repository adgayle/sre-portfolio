package main

import (
	"bytes"
	"encoding/json"
	"io"
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
	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, res, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAddBook(t *testing.T) {
	r := SetUpRouter()
	r.POST("/books", AddBook)

	// Add a valid book
	goodBook := Book{
		ISBN:     "074005352",
		Title:    "As You Like It",
		Author:   "William Shakespeare",
		Quantity: 3,
	}
	jsonValue, _ := json.Marshal(goodBook)
	goodReq, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, goodReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Attempt to add an invalid book
	badBook := Book{
		ISBN:  "222222222",
		Title: "The Great Gatsby",
	}
	jsonValue, _ = json.Marshal(badBook)
	badReq, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, badReq)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteBook(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/books", DeleteBook)

	// Book will be found.
	goodReq, _ := http.NewRequest("DELETE", "/books?isbn=183045751", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, goodReq)

	assert.Equal(t, http.StatusOK, w.Code)

	// Book will not be found.
	badReq, _ := http.NewRequest("DELETE", "/books?isbn=999999999", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, badReq)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetBooks(t *testing.T) {
	r := SetUpRouter()
	r.GET("/books", GetBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &books)
	assert.Equal(t, err, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
}

func TestCheckoutBook(t *testing.T) {
	r := SetUpRouter()

	// Book will be found.
	r.PATCH("/books", CheckoutBook)
	reqFound, _ := http.NewRequest("PATCH", "/books?isbn=123456789", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Equal(t, http.StatusOK, w.Code)

	// Book will not be found.
	reqNotFound, _ := http.NewRequest("PATCH", "/books?isbn=000000000", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Bad request missing isbn.
	reqBadNoISBN, _ := http.NewRequest("PATCH", "/books", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqBadNoISBN)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Bad request book not available.
	reqBadNotAvailable, _ := http.NewRequest("PATCH", "/books?isbn=111111111", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqBadNotAvailable)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReturnBook(t *testing.T) {
	r := SetUpRouter()
	r.PATCH("/books", ReturnBook)

	// Book will be found.
	reqFound, _ := http.NewRequest("PATCH", "/books?isbn=123456789", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Equal(t, http.StatusOK, w.Code)

	// Book not found.
	reqNotFound, _ := http.NewRequest("PATCH", "/books?isbn=000000000", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Bad return no isbn.
	reqBadNoISBN, _ := http.NewRequest("PATCH", "/books", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqBadNoISBN)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBookByISBN(t *testing.T) {
	r := SetUpRouter()
	r.GET("/books/:isbn", GetBookByISBN)

	// Book will be found.
	goodReq, _ := http.NewRequest("GET", "/books/123456789", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, goodReq)

	assert.Equal(t, http.StatusOK, w.Code)

	// Book will not be found.
	badReq, _ := http.NewRequest("GET", "/books/000000000", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, badReq)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBookByISBN(t *testing.T) {
	bookFound, _ := bookByISBN("123456789")
	assert.Equal(t, bookFound.ISBN, "123456789")

	isbn := "0000000000"
	_, err := bookByISBN(isbn)
	assert.Error(t, err, "book with ISBN "+isbn+" not found")
}
