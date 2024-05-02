package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// indicamos como se va devolver en el json
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// creamos un slice
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	//responde al cliente con los datos serializados como un json
	/**
	El status lo vamos a responder con constantes https://pkg.go.dev/net/http#pkg-constants
	*/
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	//creamos una instancia de tipo album
	var newAlbum album

	/**
	BindJSON => recibe una estructura y devuelve un error
	en este caso vamos a devolver un nuevo registro de tipo album
	*/

	//validamos el error
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func main() {
	//instancia de GIN -> trae las configuraciones por defecto
	router := gin.Default()

	//Creamos una ruta (la ruta solicita la url y el handler (controlador)) utilizando el context de gin, la ruta devuelve un JSON(por default el json solicita el status y el contenido)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})

	router.GET("/albums", getAlbums)
	//Ruta con parametro
	router.GET("/album/:id", getAlbumById)
	router.POST("/register", postAlbums)

	//Levantamos el servidor -> go run .
	router.Run("localhost:8080") //le indicamos el puerto
}
