package main

import (
    "github.com/gin-gonic/gin"
    "github.com/g0pinath/go-reference-examples/web-service-gin/controller"
)



func main() {
    router := gin.Default()
    router.GET("/albums", controller.GetAlbums)
    router.GET("/albums/:id", Controller.GetAlbumByID)
    router.POST("/albums", Controller.PostAlbums)
    router.DELETE("/albums/:id", DeleteAlbumByID)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
//     c.IndentedJSON(http.StatusOK, albums)
// }
