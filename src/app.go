package src

import (
	"go-rest-api-boilerplate/src/commons/routers"
	"log"
)

func App() {
	router := routers.Router()

	//router.Run(":8080")
	//gin.Logger("%v", router.Run("8080"))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
