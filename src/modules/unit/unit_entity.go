package unit

import "time"

type Unit struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100); not null"`
	Description *string `gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
