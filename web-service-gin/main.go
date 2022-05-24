package main

import (
    "github.com/gin-gonic/gin"
    "github.com/g0pinath/go-reference-examples/web-service-gin/controller"
)

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
    router.GET("/albums", controller.GetAlbums(albums))
    router.GET("/albums/:id", controller.GetAlbumByID(albums))
    router.POST("/albums", controller.PostAlbums(albums))
    router.DELETE("/albums/:id", controller.DeleteAlbumByID(albums))

    router.Run("localhost:8080")
}
