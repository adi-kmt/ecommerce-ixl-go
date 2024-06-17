package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func CommonHandlers(router fiber.Router, service *user_services.UserService) {
	router.Post("/login", LoginController(service))
	router.Post("/register", SignupController(service))
}
