package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func SearchProductController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productQuery := c.Query("product")
		categoryQuery := c.Query("category")
		categories := utils.StringConvertToListOfString(categoryQuery)
		productList, err := service.SearchProducts(c, productQuery, categories)
		if err != nil {
			return c.Status(err.Code).SendString(err.Message)
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(productList))
	}
}
