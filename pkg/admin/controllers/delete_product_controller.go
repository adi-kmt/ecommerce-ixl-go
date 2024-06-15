package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

func DeleteProductByIdController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productQueryParams := c.QueryInt("product_id")

		err0 := service.DeleteProduct(c, int64(productQueryParams))

		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}

		return c.Status(fiber.StatusOK).SendString("Product Deleted!!")
	}
}
