package main

import (
	"github.com/NeoListen/magic-stream-movies/server/magic-stream-movies-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
