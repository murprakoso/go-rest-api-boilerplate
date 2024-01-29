package database

import (
	"fmt"
	"go-rest-api-boilerplate/src/commons/core"
	"go-rest-api-boilerplate/src/modules/product"
	"go-rest-api-boilerplate/src/modules/unit"
	"log"
)

// Seeder runs the seeder for all tables
func Seeder() {
	fmt.Println("Running seeder")

	ProductSeeder()
}

// AutoMigration performs auto-migration for all models
func AutoMigration() {
	entities := []interface{}{
		&product.Product{},
		&unit.Unit{},
		// Add other entities here if any
	}

	err := core.DB.AutoMigrate(entities...)
	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}
}
