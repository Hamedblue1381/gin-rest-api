package main

import (
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Test1", Artist: "TestSinger1", Price: 10.99},
	{ID: "2", Title: "Test2", Artist: "TestSinger2", Price: 20.99},
	{ID: "3", Title: "Test3", Artist: "TestSinger3", Price: 30.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {

}

func postAlbums(c *gin.Context) {

}

func getAlbumByID(c *gin.Context) {

}
