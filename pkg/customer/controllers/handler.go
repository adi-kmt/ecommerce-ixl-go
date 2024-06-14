package customer_controllers

import "github.com/gofiber/fiber/v2"

func CustomerHandlers(router fiber.Router) {
	router.Post("/login")
	router.Post("/register")

	router.Get("/products")
	router.Get("/products/search")

	router.Get("/cart")
	router.Post("/cart")

	router.Post("/checkout")
}
