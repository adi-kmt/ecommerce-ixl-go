package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	admin_repositories "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/repositories"
)

type AdminService struct {
	repo *admin_repositories.AdminRepository
}

func NewAdminService(repo *admin_repositories.AdminRepository) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (service *AdminService) AddProduct(ctx *fiber.Ctx, name string, price float64, categoryID, stock int16) *messages.AppError {
	return service.repo.AddProduct(ctx, name, price, categoryID, stock)
}

func (service *AdminService) AddCategory(ctx *fiber.Ctx, name string) *messages.AppError {
	return service.repo.AddCategory(ctx, name)
}

func (service *AdminService) DeleteProduct(ctx *fiber.Ctx, id uuid.UUID) *messages.AppError {
	return service.repo.DeleteProduct(ctx, id)
}

func (service *AdminService) GetAllOrders(ctx *fiber.Ctx, userId string, status string) *messages.AppError {
	return service.repo.GetAllOrders(ctx, userId, status)
}

func (service *AdminService) ChangeOrderStatus(ctx *fiber.Ctx, orderId uuid.UUID, status string) *messages.AppError {
	return service.repo.ChangeOrderStatus(ctx, orderId, status)
}