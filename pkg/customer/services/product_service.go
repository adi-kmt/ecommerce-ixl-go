package user_services

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

func (service *UserService) GetAllProducts(ctx *fiber.Ctx) ([]entities.ProductDto, *messages.AppError) {
	return service.repo.GetProducts(ctx)
}

func (service *UserService) GetProductDetails(ctx *fiber.Ctx, id int64) (*entities.ProductDto, *messages.AppError) {
	return service.repo.GetProductDetails(ctx, id)
}

func (service *UserService) SearchProducts(ctx *fiber.Ctx, name string, category []string) ([]entities.ProductDto, *messages.AppError) {
	return service.repo.SearchProducts(ctx, name, category)
}
