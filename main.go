package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type album  struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

//slice of albums
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//-------conrtroller functions-----------
func getAlbums(c *gin.Context){

	//se devuelve el json con el slice de albums
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album


	//primero se bindea el json y luego se agrega al slice //BINDEAR es convertir un json a un objeto
	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	//si NO hay error se agrega al slice
	albums = append(albums, newAlbum)

	//se devuelve el json con el nuevo album
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context){
	id := c.Param("id")

	//se recorre el slice de albums
	for _, a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	//si no se encuentra el album se devuelve un 404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbums(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album

	// Bind the JSON to the updatedAlbum variable
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	// Find the album by ID
	for i, a := range albums {
		if a.ID == id {
			// Update the fields if provided in the request
			if updatedAlbum.Title != "" {
				albums[i].Title = updatedAlbum.Title
			}
			if updatedAlbum.Artist != "" {
				albums[i].Artist = updatedAlbum.Artist
			}
			if updatedAlbum.Price != 0 {
				albums[i].Price = updatedAlbum.Price
			}

			// Return the updated album
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}

	// If the album with the given ID is not found, return a 404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}



func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//GET ALBUMS
	router.GET("/albums", getAlbums)
	//POST ALBUMS
	router.POST("/albums", postAlbums)
	//GET ALBUMS BY ID
	router.GET("/albums/:id", getAlbumsByID)
	//UPDATE ALBUMS
	router.PUT("/albums/:id", updateAlbums)

	//------------se levanta el servidor----------------
	router.Run("localhost:8080")

}