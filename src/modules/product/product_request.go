package product

import "encoding/json"

type SProductRequest struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Qty         json.Number `json:"qty" binding:"required,number"`
	Price       json.Number `json:"price" binding:"required,number"`
}
