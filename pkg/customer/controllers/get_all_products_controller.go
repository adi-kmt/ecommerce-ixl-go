package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func GetAllProductsController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := service.GetAllProducts(c)
		if err != nil {
			return c.Status(err.Code).SendString(err.Message)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
