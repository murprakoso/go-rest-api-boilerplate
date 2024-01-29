package unit

import "gorm.io/gorm"

type IUnitRepository interface {
	FindAll() ([]Unit, error)
	FindByID(ID int) (Unit, error)
	Create(unit Unit) (Unit, error)
	Update(unit Unit) (Unit, error)
	Destroy(unit Unit) (Unit, error)
}

type SUnitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) *SUnitRepository {
	return &SUnitRepository{db}
}

func (r *SUnitRepository) FindAll() ([]Unit, error) {
	var units []Unit
	err := r.db.Find(&units).Error
	return units, err
}

func (r *SUnitRepository) FindByID(ID int) (Unit, error) {
	var unit Unit
	err := r.db.First(&unit, ID).Error
	return unit, err
}

func (r *SUnitRepository) Create(unit Unit) (Unit, error) {
	err := r.db.Create(&unit).Error
	return unit, err
}

func (r *SUnitRepository) Update(unit Unit) (Unit, error) {
	err := r.db.Save(&unit).Error
	return unit, err
}

func (r *SUnitRepository) Destroy(unit Unit) (Unit, error) {
	err := r.db.Delete(&unit).Error
	return unit, err
}
