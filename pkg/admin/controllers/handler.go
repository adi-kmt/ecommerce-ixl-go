package admin_controllers

import "github.com/gofiber/fiber/v2"

func AdminHandlers(router fiber.Router) {
	router.Patch("/admin/products")
	router.Post("/admin/products")
	router.Delete("/admin/products")

	router.Get("/admin/orders")
	router.Patch("/admin/orders")
}
