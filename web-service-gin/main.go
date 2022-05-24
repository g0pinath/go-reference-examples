package main

import (
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
)
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
    router.Run("localhost:" + httpPort)
}

// https://go.dev/doc/tutorial/web-service-gin
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. 
// (Despite the similar name, this is different from Go’s built-in context package.)
    c.IndentedJSON(http.StatusOK, albums)
//Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
}