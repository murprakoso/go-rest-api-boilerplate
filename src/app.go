package src

import (
	"go-rest-api-boilerplate/src/commons/routers"
	"log"
)

// App initializes and runs the application.
func App() {
	// Initialize the router
	router := routers.Router()

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
