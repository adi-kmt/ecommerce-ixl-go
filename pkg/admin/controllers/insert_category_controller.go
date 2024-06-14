package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

type insertCategoryDto struct {
	CategoryName string `json:"category_name"`
}

func InsertCategoryController(service *admin_services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestParams := new(insertCategoryDto)

		if err := c.BodyParser(requestParams); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}

		err0 := service.AddCategory(c, requestParams.CategoryName)

		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}

		return c.Status(fiber.StatusCreated).SendString("Category Added!!")
	}
}
