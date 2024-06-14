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
	CategoryID  int16   `json:"category_id"`
}

func InsertProductController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestParams := new(insertProductDto)

		if err := c.BodyParser(requestParams); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}

		err0 := service.AddProduct(c, requestParams.ProductName, requestParams.Price, requestParams.CategoryID, requestParams.Stock)

		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}

		return c.Status(fiber.StatusCreated).SendString("Product Added!!")
	}
}
