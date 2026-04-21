package main

import (
	"google_auth/config"
	"google_auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.EnvConfig()
	config.SetupGoogleAuth()
	config.DatabaseConfig()
	config.MigrationConfig()

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run()
}
