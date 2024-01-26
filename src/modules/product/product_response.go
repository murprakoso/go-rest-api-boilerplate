package product

import (
	"encoding/json"
	"strconv"
	"time"
)

type SProductResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Qty         json.Number `json:"qty"`
	Price       json.Number `json:"price"`
}

type SProductDetailResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Qty         json.Number `json:"qty"`
	Price       json.Number `json:"price"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"UpdatedAt"`
}

func NewProductListResponseFromEntity(products []Product) []SProductResponse {
	var productList []SProductResponse

	for _, product := range products {
		productList = append(productList, SProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Qty:         json.Number(strconv.Itoa(product.Qty)),
			Price:       json.Number(strconv.Itoa(product.Price)),
		})
	}

	return productList
}

func NewProductDetailResponseFromEntity(product Product) SProductDetailResponse {
	return SProductDetailResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Qty:         json.Number(strconv.Itoa(product.Qty)),
		Price:       json.Number(strconv.Itoa(product.Price)),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
