package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

// albums to seed record album data.
var albums = []*album{
	{ID: "48590", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "48583", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "48581", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", addAlbum)

	router.Run(":8080")
}

// getAllAlbums returns the list of all albums as JSON.
func getAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// addAlbum adds an album from JSON received in the request body.
func addAlbum(c *gin.Context) {
	var a album

	// Call ShouldBindJSON to confirm that the
	// request body JSON is valid for the struct.
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, &a)
	// Return the slice as JSON.
	c.JSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
