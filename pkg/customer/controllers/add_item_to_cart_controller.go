package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

type orderItemDto struct {
	ProductId int64  `json:"product_id"`
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

		userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64)

		orderString, err2 := service.InsertOrderItem(c, requestParams.OrderId, requestParams.ProductId, int64(userId), requestParams.Quantity)

		if err2 != nil {
			return c.Status(err2.Code).SendString(err2.Message)
		}
		if orderString == "" {
			return c.Status(fiber.StatusCreated).SendString("Order Item created")
		} else {
			return c.Status(fiber.StatusCreated).SendString("Order Id is " + orderString)
		}
	}
}
