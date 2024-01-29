package unit

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-boilerplate/src/commons/core"
)

var unitHandler *SUnitHandler

func InitUnitModule() {
	unitRepository := NewUnitRepository(core.DB)
	unitService := NewUnitService(unitRepository)
	unitHandler = NewUnitHandler(unitService)
}

func SetUnitRouterGroup(router *gin.RouterGroup) {
	router.GET("/unit", unitHandler.ShowUnits)
	router.GET("/unit/:id", unitHandler.ShowUnit)
	router.POST("/unit", unitHandler.CreateUnit)
	router.PUT("/unit/:id", unitHandler.UpdateUnit)
	router.DELETE("/unit/:id", unitHandler.DestroyUnit)
}
