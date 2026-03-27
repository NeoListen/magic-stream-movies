package routes

import (
	controller "github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})
}
