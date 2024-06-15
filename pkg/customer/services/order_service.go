package user_services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

func (service *UserService) InsertOrderItem(ctx *fiber.Ctx, orderId string, productId int64, userId int64, quantity int16) (string, *messages.AppError) {
	if orderId == "" {
		return service.repo.InsertIntoOrderAndOrderItems(ctx, productId, userId, quantity)
	} else {
		orderUUID, err := uuid.Parse(orderId)
		if err != nil {
			log.Debugf("Error Parsing Order ID: %v", err)
			return "", messages.BadRequest("Error Parsing Order ID")
		}
		return "", service.repo.InsertItemIntoOrderItem(ctx, productId, orderUUID, userId, quantity)
	}
}

func (service *UserService) GetAllItemsInOrder(ctx *fiber.Ctx, orderId uuid.UUID) (*entities.OrderDto, *messages.AppError) {
	return service.repo.GetItemsInCart(ctx, orderId)
}

func (service *UserService) AddPaymentIdToOrder(ctx *fiber.Ctx, orderId, paymentId uuid.UUID) *messages.AppError {
	return service.repo.UpdateOrderPaymentId(ctx, orderId, paymentId)
}
