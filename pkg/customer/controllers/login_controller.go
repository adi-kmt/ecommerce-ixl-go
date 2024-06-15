package customer_controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
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
		claims := jwt.MapClaims{
			"Email": requestParams.Email,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		secretKey, isPresent := os.LookupEnv("JWT_SECRET")
		if !isPresent {
			secretKey = "sample_secret"
		}

		t, err1 := token.SignedString(secretKey)
		if err1 != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString("Something went wrong while signing the token")
		}

		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(map[string]interface{}{"token": t}))
	}
}
