package main

import (
    "github.com/gin-gonic/gin"
    "github.com/g0pinath/go-reference-examples/web-service-gin/controller"
)



func main() {
    router := gin.Default()
    router.GET("/albums", controller.GetAlbums)
    router.GET("/albums/:id", controller.GetAlbumByID)
    router.POST("/albums", controller.PostAlbums)
    router.DELETE("/albums/:id", controller.DeleteAlbumByID)

    router.Run("localhost:8080")
}
