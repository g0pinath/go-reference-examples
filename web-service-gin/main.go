package main

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/g0pinath/go-reference-examples/controllers"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", controllers.getAlbums)
    router.GET("/albums/:id", controllers.getAlbumByID)
    router.POST("/albums", controllers.postAlbums)
    router.DELETE("/albums/:id", controllers.deleteAlbumByID)

    router.Run("localhost:8080")
}


// curl "http://localhost:8080/albums/2"   --request "DELETE"