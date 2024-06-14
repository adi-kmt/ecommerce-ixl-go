package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

type CheckoutRequest struct {
	PaymentId string `json:"payment_id"`
	OrderId   string `json:"order_id"`
}

func CheckoutController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestParams := new(CheckoutRequest)
		if err := c.BodyParser(requestParams); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}
		orderUUID, err0 := uuid.Parse(requestParams.OrderId)
		if err0 != nil {
			return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Order ID")
		}
		paymentUUID, err2 := uuid.Parse(requestParams.PaymentId)
		if err2 != nil {
			return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Payment ID")
		}
		err1 := service.AddPaymentIdToOrder(c, orderUUID, paymentUUID)
		if err1 != nil {
			return c.Status(err1.Code).SendString(err1.Message)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
