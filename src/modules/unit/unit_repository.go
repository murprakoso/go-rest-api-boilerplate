package unit

import (
	"go-rest-api-boilerplate/src/commons/core"
	"gorm.io/gorm"
)

type IUnitRepository interface {
	FindAll(paginationQuery core.PaginationQuery) ([]Unit, int, error)
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

func (r *SUnitRepository) FindAll(paginationQuery core.PaginationQuery) ([]Unit, int, error) {
	var units []Unit
	var total int64
	offset := (paginationQuery.Page - 1) * paginationQuery.Limit
	query := r.db.Offset(offset).Limit(paginationQuery.Limit)

	// Add search condition if search parameter is provided
	if paginationQuery.Search != "" {
		query = query.Where("name LIKE ?", "%"+paginationQuery.Search+"%")
		// Gunakan Or untuk menambahkan kondisi pencarian pada kolom name atau description
		//query = query.Where("name LIKE ?", "%"+paginationQuery.Search+"%").
		//	Or("description LIKE ?", "%"+paginationQuery.Search+"%")
	}

	// Add order by and sort by conditions
	query = query.Order(paginationQuery.OrderBy + " " + paginationQuery.SortBy)

	// Count total number of records
	if err := r.db.Model(&Unit{}).Where("name LIKE ?", "%"+paginationQuery.Search+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// Count total number of records (Jika menggunakan kondisi Or)
	//if err := r.db.Model(&Unit{}).
	//	Where("name LIKE ?", "%"+paginationQuery.Search+"%").
	//	Or("description LIKE ?", "%"+paginationQuery.Search+"%").
	//	Count(&total).Error; err != nil {
	//	return nil, 0, err
	//}

	err := query.Find(&units).Error
	return units, int(total), err
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
