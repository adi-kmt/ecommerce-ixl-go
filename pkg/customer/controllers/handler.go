package customer_controllers

import (
	"github.com/gofiber/fiber/v2"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func CustomerHandlers(router fiber.Router, service *user_services.UserService) {
	router.Get("/profile", GetProfileController(service))

	router.Get("/products", GetAllProductsController(service))
	router.Get("/products/search", SearchProductController(service))

	router.Get("/cart", GetCartItemsController(service))
	router.Post("/cart", AddItemToController(service))

	router.Post("/checkout", CheckoutController(service))
}
