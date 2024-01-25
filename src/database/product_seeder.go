package database

import (
	"fmt"
	"go-rest-api-boilerplate/src/commons/core"
	entity "go-rest-api-boilerplate/src/modules/product"
)

func ProductSeeder() {
	fmt.Println("Running Product Seeder")
	product := entity.Product{}
	product.Name = "Laptop 21 In"
	product.Description = "The Description"
	product.Price = 50000

	err := core.DB.Create(&product).Error
	if err != nil {
		fmt.Println("Error creating product record!")
	}
}
