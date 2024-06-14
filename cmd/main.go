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

	// Monitoring the requests made to this service.
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	api := app.Group("/api/v1")

	admin_controllers.AdminHandlers(api)
	customer_controllers.CustomerHandlers(api)

	port, isErr := os.LookupEnv("API_PORT")
	if !isErr {
		port = "8080"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
