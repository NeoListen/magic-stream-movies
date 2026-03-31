package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/database"
	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to load .env file")
	}

	allowOrigin := os.Getenv("ALLOW_ORIGIN")
	var origins []string
	if allowOrigin != "" {
		origins = strings.Split(allowOrigin, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Printf("Allowed Origin %s", origins[i])
		}
	} else {
		origins = []string{"http://localhost:8080"}
		log.Printf("Allowed Origin %s", origins[0])
	}

	config := cors.Config{}
	//config.AllowOrigins = origins
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	//config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))
	router.Use(gin.Logger())

	var client *mongo.Client = database.Connect()

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}(client, context.Background())

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
