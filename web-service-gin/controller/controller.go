package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context, albums *gin.Context, []controller.album) {
    c.IndentedJSON(http.StatusOK, albums)
}

func remove(albums []album, i int) []album {
	copy(albums[i:], albums[i+1:])
	return albums[:len(albums)-1]
}

func DeleteAlbumByID(c *gin.Context, albums []album) {

    id := c.Param("id")
    index_slice := 0

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            albums = remove(albums, index_slice)
            return 
        }
        index_slice += 1
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context, albums []album) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context, albums []album) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}