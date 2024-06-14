package user_services

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
)

func (service *UserService) GetUserDetailsAndOrders(ctx *fiber.Ctx, id int64) *messages.AppError {
	return service.repo.GetUserDetailsAndOrders(ctx, id)
}

func (service *UserService) InsertUser(ctx *fiber.Ctx, name, email, address, password string, isAdmin bool) *messages.AppError {
	return service.repo.InsertUser(ctx, name, email, address, password, isAdmin)
}

func (service *UserService) ValidateUser(ctx *fiber.Ctx, email, password string) *messages.AppError {
	err := service.repo.GetUserDetails(ctx, email)
	if err != nil {
		return err
	}
	//TODO Check if password from user and DB is same else send unauthorized user
	return nil
}
