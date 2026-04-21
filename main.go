package main

import (
	"google_auth/config"
	"google_auth/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.EnvConfig()
	config.SetupGoogleAuth()
	config.DatabaseConfig()
	config.MigrationConfig()
}

func main() {
	router := gin.Default()

	routes.UserRoutes(router)

	router.Run()
}
