package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
)

func AdminHandlers(router fiber.Router, service *admin_services.AdminService) {
	router.Post("/admin/products", InsertProductController(service))
	router.Delete("/admin/products", DeleteProductByIdController(service))

	router.Post("/admin/categories", InsertCategoryController(service))

	router.Get("/admin/orders", GetOrdersController(service))
}
