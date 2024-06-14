package messages

import "github.com/gofiber/fiber/v2"

// SuccessResponseSlice is the list SuccessResponse that will be passed in the response by Handler
func SuccessResponseSlice[T any](data []T) *fiber.Map {
	return &fiber.Map{
		"data":  data,
		"error": "",
	}
}

// SuccessResponse is the primitive type SuccessResponse that will be passed in the response by Handler
func SuccessResponse[T any](data T) *fiber.Map {
	return &fiber.Map{
		"data":  data,
		"error": "",
	}
}

// ErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  "",
		"error": err.Error(),
	}
}
