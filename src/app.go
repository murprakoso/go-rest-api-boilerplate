package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	commons "go-rest-api-boilerplate/src/commons/config"
	"go-rest-api-boilerplate/src/commons/core"
	"go-rest-api-boilerplate/src/commons/routers"
	"log"
)

// App initializes and runs the application.
func App() {

	// Load configuration from .env file
	config := commons.LoadConfig()

	// Use the loaded configuration
	gin.SetMode(gin.ReleaseMode)
	fmt.Printf("Environment Mode: %s\n", config.EnvMode)
	//fmt.Printf("Port: %s\n", config.Port)
	//fmt.Printf("DB Host: %s\n", config.DBHost)
	//fmt.Printf("DB Port: %s\n", config.DBPort)
	//fmt.Printf("DB User: %s\n", config.DBUser)
	//fmt.Printf("DB Pass: %s\n", config.DBPass)
	//fmt.Printf("DB Name: %s\n", config.DBName)

	// Initialize the router
	router := routers.Router()

	// Initialize the setup CORS
	router.Use(core.SetupCors())

	// Run the server on port 8080
	err := router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// setReleaseMode sets GIN_MODE to "release" if not already set.
//func setReleaseMode() {
//	mode := os.Getenv("GIN_MODE")
//	if mode == "" {
//		os.Setenv("GIN_MODE", "release")
//	}
//}
