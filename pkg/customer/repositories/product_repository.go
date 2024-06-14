package user_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
)

func (repo *UserRepository) GetProducts(ctx *fiber.Ctx) *messages.AppError {

	_, err := repo.q.GetProductsForCategories(ctx.Context(), []string{"Featured"})
	if err != nil {
		log.Debugf("Error Getting Products: %v", err)
		return messages.InternalServerError("Error Getting Products")
	}
	return nil
}

func (repo *UserRepository) SearchProducts(ctx *fiber.Ctx, name string, category []string) *messages.AppError {

	_, err := repo.q.SearchProducts(ctx.Context(), db.SearchProductsParams{
		Column1: &name,
		Column2: category,
	})
	if err != nil {
		log.Debugf("Error Searching Products: %v", err)
		return messages.InternalServerError("Error Searching Products")
	}
	return nil
}

func (repo *UserRepository) GetProductDetails(ctx *fiber.Ctx, id uuid.UUID) *messages.AppError {

	pgUUID := utils.ConvertUUIDToPgType(id)
	_, err := repo.q.GetProductDetailByID(ctx.Context(), pgUUID)
	if err != nil {
		log.Debugf("Error Getting Product: %v", err)
		return messages.InternalServerError("Error Getting Product")
	}
	return nil
}
