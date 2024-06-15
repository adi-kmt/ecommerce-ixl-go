package entities

import (
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
)

type ProductDto struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int16   `json:"stock"`
	CategoryID  int16   `json:"category_id"`
	Description string  `json:"description"`
}

func ProductDtoFromDbRowSingle(row *db.Product) *ProductDto {
	return &ProductDto{
		ID:          utils.ConvertPgUUIDToString(row.ID),
		Name:        row.Name,
		Price:       row.Price,
		Stock:       row.Stock,
		CategoryID:  row.CategoryID,
		Description: row.Description,
	}
}

func ProductDtoFromDbRow(row []*db.Product) []ProductDto {
	var productList []ProductDto

	for _, product := range row {
		productList = append(productList, ProductDto{
			ID:          utils.ConvertPgUUIDToString(product.ID),
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			CategoryID:  product.CategoryID,
			Description: product.Description,
		})
	}
	return productList
}
