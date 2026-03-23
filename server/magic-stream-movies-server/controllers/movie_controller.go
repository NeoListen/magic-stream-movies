package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/database"
	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")
var validate = validator.New()

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var movies []models.Movie
		cursor, err := movieCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch movies"})
			return
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to close cursor"})
				return
			}
		}(cursor, ctx)
		if err = cursor.All(ctx, &movies); err != nil {
			print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode movies"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"movies": movies})
	}
}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		movieId := c.Param("imdb_id")
		if movieId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "movie id is required"})
			return
		}
		var movie models.Movie
		if err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieId}).Decode(&movie); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode movie"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"movie": movie})
	}
}

func AddMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}
		one, err := movieCollection.InsertOne(ctx, movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert movie"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"movie": one})
	}
}
