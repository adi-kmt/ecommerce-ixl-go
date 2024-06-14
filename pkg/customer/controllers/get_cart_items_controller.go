package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func GetCartItemsController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		orderId := c.Query("order_id")
		orderUUID, err := uuid.Parse(orderId)
		if err != nil {
			log.Debugf("Error Parsing Order ID: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Order ID")
		}
		err0 := service.GetAllItemsInOrder(c, orderUUID)
		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
