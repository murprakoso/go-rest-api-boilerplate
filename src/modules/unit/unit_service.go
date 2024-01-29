package unit

import (
	"go-rest-api-boilerplate/src/commons/core"
	"math"
)

type IUnitService interface {
	FindAll(paginationQuery core.PaginationQuery) ([]Unit, core.PaginationMeta, error)
	FindByID(ID int) (Unit, error)
	Create(unit SUnitRequest) (Unit, error)
	Update(ID int, unitRequest SUnitRequest) (Unit, error)
	Destroy(ID int) (Unit, error)
}

type SUnitService struct {
	unitRepository IUnitRepository
}

func NewUnitService(unitRepository IUnitRepository) *SUnitService {
	return &SUnitService{unitRepository}
}

func (s *SUnitService) FindAll(paginationQuery core.PaginationQuery) ([]Unit, core.PaginationMeta, error) {
	units, total, err := s.unitRepository.FindAll(paginationQuery)
	if err != nil {
		return nil, core.PaginationMeta{}, err
	}

	// Calculate pagination meta
	totalPages := int(math.Ceil(float64(total) / float64(paginationQuery.Limit)))
	currentPage := paginationQuery.Page
	itemCount := len(units)

	meta := core.PaginationMeta{
		TotalItems:   total,
		ItemCount:    itemCount,
		ItemsPerPage: paginationQuery.Limit,
		TotalPages:   totalPages,
		CurrentPage:  currentPage,
	}

	return units, meta, nil
}

func (s *SUnitService) FindByID(ID int) (Unit, error) {
	return s.unitRepository.FindByID(ID)
}

func (s *SUnitService) Create(unitRequest SUnitRequest) (Unit, error) {
	unit := Unit{
		Name:        unitRequest.Name,
		Description: &unitRequest.Description,
	}
	createdUnit, err := s.unitRepository.Create(unit)
	return createdUnit, err
}

func (s *SUnitService) Update(ID int, unitRequest SUnitRequest) (Unit, error) {
	unit, err := s.unitRepository.FindByID(ID)
	if err != nil {
		return Unit{}, err
	}
	unit.Name = unitRequest.Name
	unit.Description = &unitRequest.Description

	updatedUnit, err := s.unitRepository.Update(unit)
	return updatedUnit, err
}

func (s *SUnitService) Destroy(ID int) (Unit, error) {
	unit, err := s.unitRepository.FindByID(ID)
	deletedUnit, err := s.unitRepository.Destroy(unit)
	return deletedUnit, err
}
