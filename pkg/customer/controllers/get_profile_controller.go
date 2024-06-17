package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func GetProfileController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(float64)
		profile, err := service.GetUserDetailsAndOrders(c, int64(userId))
		if err != nil {
			return c.Status(err.Code).SendString(err.Message)
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(profile))
	}
}
