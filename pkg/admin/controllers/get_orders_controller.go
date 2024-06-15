package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

func GetOrdersController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Query("user_id")
		status := c.Query("status")
		orders, err := service.GetAllOrders(c, userId, status)
		if err != nil {
			return c.Status(err.Code).SendString(err.Message)
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(orders))
	}
}
