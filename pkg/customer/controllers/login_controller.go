package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

type loginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestParams := new(loginDto)
		if err := c.BodyParser(requestParams); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}
		err0 := service.ValidateUser(c, requestParams.Email, requestParams.Password)
		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}
		return c.Status(fiber.StatusOK).SendString("Login Successful!!")
	}
}
