package config

import "google_auth/models"

func MigrationConfig() {
	DB.AutoMigrate(&models.User{})
}
