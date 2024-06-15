package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

type insertProductDto struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Stock       int16   `json:"stock"`
	CategoryID  int32   `json:"category_id"`
}

func InsertProductController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(insertProductDto)

		if err := c.BodyParser(requestBody); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}

		err0 := service.AddProduct(c, requestBody.ProductName, requestBody.Price, requestBody.CategoryID, requestBody.Stock)

		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}

		return c.Status(fiber.StatusCreated).SendString("Product Added!!")
	}
}
