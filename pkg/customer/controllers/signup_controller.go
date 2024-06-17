package customer_controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
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
		userId, err0 := service.InsertUser(c, requestParams.Name, requestParams.Email, requestParams.Address, requestParams.Password, requestParams.IsAdmin)
		if err0 != nil {
			return c.Status(err0.Code).SendString(err0.Message)
		}
		claims := jwt.MapClaims{
			"email": requestParams.Email,
			"id":    userId,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		secretKey, isPresent := os.LookupEnv("JWT_SECRET")
		if !isPresent {
			secretKey = "sample_secret"
		}

		t, err1 := token.SignedString([]byte(secretKey))
		if err1 != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString("Something went wrong while signing the token")
		}
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse(map[string]interface{}{"token": t}))
	}
}
