package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

type orderItemDto struct {
	UserId    int64  `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int16  `json:"quantity"`
	OrderId   string `json:"order_id"`
}

func AddItemToController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestParams := new(orderItemDto)
		if err := c.BodyParser(requestParams); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}

		productUUID, err1 := uuid.Parse(requestParams.ProductId)
		if err1 != nil {
			log.Debugf("Error Parsing Product ID: %v", err1)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Product ID")
		}
		orderString, err2 := service.InsertOrderItem(c, requestParams.OrderId, productUUID, requestParams.UserId, requestParams.Quantity)

		if err2 != nil {
			return c.Status(err2.Code).SendString(err2.Message)
		}
		if orderString != "" {
			return c.Status(fiber.StatusCreated).SendString("Order Item created")
		} else {
			return c.Status(fiber.StatusCreated).SendString("Order Id is " + orderString)
		}
	}
}
