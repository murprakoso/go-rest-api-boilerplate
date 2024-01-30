package auth

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Username  string `gorm:"type:varchar(100); not null; unique"`
	Email     string `gorm:"type:varchar(100); unique"`
	Password  string `gorm:"type:varchar(200)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
