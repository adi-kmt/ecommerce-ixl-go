package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

func DeleteProductByIdController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productQueryParams := c.Query("product_id")

		productId, err := uuid.Parse(productQueryParams)
		if err != nil {
			log.Debugf("Error Parsing Product ID: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Product ID")
		}

		err0 := service.DeleteProduct(c, productId)

		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}

		return c.Status(fiber.StatusOK).SendString("Product Deleted!!")
	}
}
