package routes

import (
	controller "github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/controllers"
	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.POST("/movie", controller.AddMovie(client))
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate(client))
	router.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
}
