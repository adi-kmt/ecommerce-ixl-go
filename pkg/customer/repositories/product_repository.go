package user_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

func (repo *UserRepository) GetProducts(ctx *fiber.Ctx) ([]entities.ProductDto, *messages.AppError) {

	products, err := repo.q.GetProductsForCategories(ctx.Context(), []int32{1})
	if err != nil {
		log.Debugf("Error Getting Products: %v", err)
		return nil, messages.InternalServerError("Error Getting Products")
	}
	return entities.ProductDtoFromDbRow(products), nil
}

func (repo *UserRepository) SearchProducts(ctx *fiber.Ctx, name string, category []int32) ([]entities.ProductDto, *messages.AppError) {

	products, err := repo.q.SearchProducts(ctx.Context(), db.SearchProductsParams{
		Column1: &name,
		Column2: category,
	})
	if err != nil {
		log.Debugf("Error Searching Products: %v", err)
		return nil, messages.InternalServerError("Error Searching Products")
	}
	return entities.ProductDtoFromDbRow(products), nil
}

func (repo *UserRepository) GetProductDetails(ctx *fiber.Ctx, id int64) (*entities.ProductDto, *messages.AppError) {

	product, err := repo.q.GetProductDetailByID(ctx.Context(), id)
	if err != nil {
		log.Debugf("Error Getting Product: %v", err)
		return nil, messages.InternalServerError("Error Getting Product")
	}
	return entities.ProductDtoFromDbRowSingle(product), nil
}
