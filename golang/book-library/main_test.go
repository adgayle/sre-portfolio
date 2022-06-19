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
	"github.com/go-playground/assert/v2"
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
