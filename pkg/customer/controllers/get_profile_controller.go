package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func GetProfileController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.QueryInt("id")
		profile, err := service.GetUserDetailsAndOrders(c, int64(id))
		if err != nil {
			return c.Status(err.Code).SendString(err.Message)
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(profile))
	}
}
