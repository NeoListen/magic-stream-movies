package main

import (
	controller "github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/movie", controller.AddMovie())

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
