package user_services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
)

func (service *UserService) GetAllProducts(ctx *fiber.Ctx, id int64) *messages.AppError {
	return service.repo.GetProducts(ctx)
}

func (service *UserService) GetProductDetails(ctx *fiber.Ctx, id uuid.UUID) *messages.AppError {
	return service.repo.GetProductDetails(ctx, id)
}

func (service *UserService) SearchProducts(ctx *fiber.Ctx, name string, category []string) *messages.AppError {
	return service.repo.SearchProducts(ctx, name, category)
}
