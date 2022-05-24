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


func main() {
    router := gin.Default()
    router.GET("/albums", controller.GetAlbums)
    // router.GET("/albums/:id", controller.GetAlbumByID(albums))
    // router.POST("/albums", controller.PostAlbums(albums))
    // router.DELETE("/albums/:id", controller.DeleteAlbumByID(albums))

    router.Run("localhost:8080")
}
