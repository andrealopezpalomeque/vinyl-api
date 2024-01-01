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

	router.Run("localhost:8080")

}