package unit

import (
	"time"
)

type SUnitResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SUnitDetailResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

func NewUnitListResponseFromEntity(units []Unit) []SUnitResponse {
	var unitList []SUnitResponse

	for _, unit := range units {
		unitList = append(unitList, SUnitResponse{
			ID:          unit.ID,
			Name:        unit.Name,
			Description: *unit.Description,
		})
	}

	return unitList
}

func NewUnitDetailResponseFromEntity(unit Unit) SUnitDetailResponse {
	return SUnitDetailResponse{
		ID:          unit.ID,
		Name:        unit.Name,
		Description: *unit.Description,
		CreatedAt:   unit.CreatedAt,
		UpdatedAt:   unit.UpdatedAt,
	}
}
