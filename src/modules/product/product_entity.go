package product

import "time"

// Product represents the data model for a product (product entity).
type Product struct {
	ID          int
	Name        string
	Description string
	Qty         int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
