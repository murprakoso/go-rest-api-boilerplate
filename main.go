package main

import (
	"go-rest-api-boilerplate/src"
	"go-rest-api-boilerplate/src/commons/core"
	"go-rest-api-boilerplate/src/database"
)

func main() {
	// Initialize the database connection
	core.InitializeDatabase()

	// Perform auto-migration for database models
	database.AutoMigration()

	// Run the seeder to populate the database with initial data
	database.Seeder()

	// Start the application
	src.App()
}
