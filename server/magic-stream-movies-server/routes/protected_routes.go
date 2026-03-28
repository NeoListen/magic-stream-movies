package routes

import (
	controller "github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/controllers"
	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/movie", controller.AddMovie())
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate())
	router.GET("/recommendedmovies", controller.GetRecommendedMovies())
}
