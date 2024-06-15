package customer_controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func GetAllProductsController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryProductId := c.Query("product_id")
		if queryProductId == "" {
			productList, err := service.GetAllProducts(c)
			if err != nil {
				return c.Status(err.Code).SendString(err.Message)
			}
			return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(productList))
		} else {
			productValue, err0 := strconv.Atoi(queryProductId)
			if err0 != nil {
				return c.Status(fiber.ErrBadRequest.Code).SendString("Product ID is not a number")
			}
			product, err := service.GetProductDetails(c, int64(productValue))
			if err != nil {
				return c.Status(err.Code).SendString(err.Message)
			}
			return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(product))
		}
	}
}
