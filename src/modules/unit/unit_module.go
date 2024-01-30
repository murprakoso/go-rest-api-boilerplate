package unit

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var unitHandler *SUnitHandler

func InitRouterGroup(router *gin.RouterGroup, db *gorm.DB) {
	unitRepository := NewUnitRepository(db)
	unitService := NewUnitService(unitRepository)
	unitHandler = NewUnitHandler(unitService)

	router.GET("/unit", unitHandler.ShowUnits)
	router.GET("/unit/:id", unitHandler.ShowUnit)
	router.POST("/unit", unitHandler.CreateUnit)
	router.PUT("/unit/:id", unitHandler.UpdateUnit)
	router.DELETE("/unit/:id", unitHandler.DestroyUnit)
}
