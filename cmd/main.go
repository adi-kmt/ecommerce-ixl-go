package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/injection"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/jwt"
	admin_controllers "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/controllers"
	customer_controllers "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/controllers"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Main File %s", err.Error())
	}

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	secretKey, isPresent := os.LookupEnv("JWT_SECRET")
	if !isPresent {
		secretKey = "sample_secret"
	}

	userService, adminService := injection.InjectDependencies()

	api := app.Group("/api/v1")

	customer_controllers.CommonHandlers(api, userService)

	app.Use(jwt.NewAuthMiddleware(secretKey))

	// Monitoring the requests made to this service.
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	admin_controllers.AdminHandlers(api, adminService)
	customer_controllers.CustomerHandlers(api, userService)

	port, isErr := os.LookupEnv("API_PORT")
	if !isErr {
		port = "8080"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
