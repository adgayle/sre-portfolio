package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// main the controlling function.
func main() {
	apiAddress := flag.String("address", "", "Address running the API.")
	apiPort := flag.String("port", "8080", "Book API HTTP port.")
	releaseMode := flag.Bool("release-mode", true, "Gin execution mode true=release|false=debug.")

	if *releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to configure trusted proxies: %v", err)
	}

	router.GET("/", HandleHomePage)
	router.GET("/books", GetBooks)
	router.POST("/books", AddBook)
	router.GET("/books/:isbn", GetBookByISBN)
	router.PATCH("/checkout", CheckoutBook)
	router.PATCH("/return", ReturnBook)
	router.DELETE("/books/:isbn", DeleteBook)
	router.GET("/healthz", healthCheck)
	router.GET("/readyz", readyCheck)

	addr := fmt.Sprintf("%v:%v", *apiAddress, *apiPort)
	err = router.Run(addr)
	if err != nil {
		log.Fatalf("Failed to run router on %v with error %v", addr, err)
	}

	// TODO: add a proper shutdown.
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Status OK"})
}

func readyCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Status OK"})
}
