package auth

import "gorm.io/gorm"

type IAuthRepository interface {
	Create(user User) (User, error)
	FindByUsername(username string) (User, error)
	FindByEmail(email string) (User, error)
}

type SAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *SAuthRepository {
	return &SAuthRepository{db}
}

func (r *SAuthRepository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *SAuthRepository) FindByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *SAuthRepository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
