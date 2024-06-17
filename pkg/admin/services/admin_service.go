package admin_services

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	admin_repositories "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/repositories"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

type AdminService struct {
	repo *admin_repositories.AdminRepository
}

func NewAdminService(repo *admin_repositories.AdminRepository) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (service *AdminService) AddProduct(ctx *fiber.Ctx, name string, price float64, categoryID int32, stock int16) *messages.AppError {
	return service.repo.AddProduct(ctx, name, price, categoryID, stock)
}

func (service *AdminService) AddCategory(ctx *fiber.Ctx, name string) *messages.AppError {
	return service.repo.AddCategory(ctx, name)
}

func (service *AdminService) DeleteProduct(ctx *fiber.Ctx, id int64) *messages.AppError {
	return service.repo.DeleteProduct(ctx, id)
}

func (service *AdminService) GetAllOrders(ctx *fiber.Ctx, userId string, status string) ([]entities.AdminOrderDto, *messages.AppError) {
	return service.repo.GetAllOrders(ctx, userId, status)
}
