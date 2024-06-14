package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

type signupDto struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Address       string `json:"address"`
	Name          string `json:"name"`
	IsAdmin       bool   `json:"is_admin"`
	AdminPassword string `json:"admin_password"`
}

func SignupController(service *user_services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestParams := new(signupDto)
		if err := c.BodyParser(requestParams); err != nil {
			log.Debugf("Error parsing request body: %v", err)
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error parsing request body")
		}
		if requestParams.IsAdmin {
			if requestParams.AdminPassword != "admin" {
				return c.Status(fiber.StatusUnauthorized).SendString("Invalid Admin Password")
			}
		}
		err0 := service.InsertUser(c, requestParams.Name, requestParams.Email, requestParams.Address, requestParams.Password, requestParams.IsAdmin)
		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}
		return c.Status(fiber.StatusCreated).SendString("Signup Successful!!")
	}
}
