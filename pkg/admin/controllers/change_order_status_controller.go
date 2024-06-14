package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

func PatchOrderStatusController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		orderId := c.Query("order_id")
		status := c.Query("status")

		orderUUID, err := uuid.Parse(orderId)
		if err != nil {
			log.Debugf("Error Parsing Order ID: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Order ID")
		}
		err0 := service.ChangeOrderStatus(c, orderUUID, status)
		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}
		return c.Status(fiber.StatusOK).SendString("Order Status Updated!!")
	}
}
