package user_services

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

func (service *UserService) GetUserDetailsAndOrders(ctx *fiber.Ctx, id int64) (*entities.UserDetailsAndOrdersDto, *messages.AppError) {
	return service.repo.GetUserDetailsAndOrders(ctx, id)
}

func (service *UserService) InsertUser(ctx *fiber.Ctx, name, email, address, password string, isAdmin bool) (int64, *messages.AppError) {
	return service.repo.InsertUser(ctx, name, email, address, password, isAdmin)
}

func (service *UserService) ValidateUser(ctx *fiber.Ctx, email, password string) (int64, *messages.AppError) {
	user, err := service.repo.GetUserDetails(ctx, email)
	if err != nil {
		return 0, err
	}
	if user.Password == password {
		return user.ID, nil
	} else {
		return 0, messages.Unauthorized("Invalid Password")
	}
}
