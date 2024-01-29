package unit

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-boilerplate/src/commons/core"
	"net/http"
	"strconv"
)

type SUnitHandler struct {
	unitService IUnitService
}

func NewUnitHandler(unitService IUnitService) *SUnitHandler {
	return &SUnitHandler{unitService}
}

// ShowUnits handles the HTTP GET request to retrieve all units.
func (h *SUnitHandler) ShowUnits(c *gin.Context) {
	paginationQuery := core.ParsePaginationQuery(c)

	units, meta, err := h.unitService.FindAll(paginationQuery)
	if err != nil {
		core.ResponseError(c, http.StatusInternalServerError, "Invalid request payload")
		return
	}

	core.ResponsePaginate(c, http.StatusOK, true, "", NewUnitListResponseFromEntity(units), meta)
}

// ShowUnit handles the HTTP GET request to retrieve a single unit by ID.
func (h *SUnitHandler) ShowUnit(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	unit, err := h.unitService.FindByID(ID)
	if err != nil {
		core.ResponseJSON(c, http.StatusNotFound, false, "Unit not found", nil)
		return
	}

	core.ResponseJSON(c, http.StatusOK, true, "", NewUnitDetailResponseFromEntity(unit))
}

// CreateUnit handles the HTTP POST request to create a new unit.
func (h *SUnitHandler) CreateUnit(c *gin.Context) {
	var unitRequest SUnitRequest
	if err := c.ShouldBindJSON(&unitRequest); err != nil {
		core.ResponseError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdUnit, err := h.unitService.Create(unitRequest)
	if err != nil {
		core.ResponseJSON(c, http.StatusBadRequest, false, "Failed to create unit", nil)
		return
	}

	core.ResponseJSON(c, http.StatusCreated, true, "", NewUnitDetailResponseFromEntity(createdUnit))
}

// UpdateUnit handles the request to update a unit.
func (h *SUnitHandler) UpdateUnit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	var unitRequest SUnitRequest
	if err := c.ShouldBindJSON(&unitRequest); err != nil {
		core.ResponseError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedUnit, err := h.unitService.Update(ID, unitRequest)
	if err != nil {
		core.ResponseError(c, http.StatusBadRequest, "Failed to update unit")
		return
	}

	core.ResponseJSON(c, http.StatusOK, true, "Update data successfully", NewUnitDetailResponseFromEntity(updatedUnit))
}

// DestroyUnit remove a data record
func (h *SUnitHandler) DestroyUnit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	deletedUnit, err := h.unitService.Destroy(ID)
	if err != nil {
		core.ResponseJSON(c, http.StatusBadRequest, false, "Failed to delete unit", nil)
		return
	}

	core.ResponseJSON(c, http.StatusOK, true, "", NewUnitDetailResponseFromEntity(deletedUnit))
}
