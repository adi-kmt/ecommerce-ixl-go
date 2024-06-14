package user_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
)

func (repo *UserRepository) GetItemsInCart(ctx *fiber.Ctx, orderId uuid.UUID) *messages.AppError {

	pgUUID := utils.ConvertUUIDToPgType(orderId)
	_, err := repo.q.GetCurrentOrderByID(ctx.Context(), pgUUID)
	if err != nil {
		log.Debugf("Error Getting Items In Cart: %v", err)
		return messages.InternalServerError("Error Getting Items In Cart")
	}
	return nil
}

func (repo *UserRepository) InsertItemIntoOrderItem(ctx *fiber.Ctx, orderId, productId uuid.UUID, userId int64, quantity int16) *messages.AppError {

	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)
	pgProductUUID := utils.ConvertUUIDToPgType(productId)

	product, err0 := repo.q.GetProductDetailByID(ctx.Context(), pgProductUUID)
	if err0 != nil {
		log.Debugf("Error Getting Product: %v", err0)
		return messages.InternalServerError("Error Getting Product")
	}

	priceAgg := product.Price * float64(quantity)
	err1 := repo.q.InsertIntoOrderItemsTable(ctx.Context(), db.InsertIntoOrderItemsTableParams{
		OrderID:         pgOrderUUID,
		ProductID:       pgProductUUID,
		UserID:          userId,
		ProductQuantity: quantity,
		ProductPriceAgg: priceAgg,
	})
	if err1 != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err1)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}
	return nil
}

func (repo *UserRepository) InsertIntoOrderAndOrderItems(ctx *fiber.Ctx, productId uuid.UUID, userId int64, quantity int16) *messages.AppError {
	orderId, err := utils.GenerateNewUUID()
	if err != nil {
		return messages.InternalServerError("Error Generating UUID")
	}
	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)
	pgProductUUID := utils.ConvertUUIDToPgType(productId)
	paymentUUID := utils.ConvertUUIDToPgType(uuid.Nil)
	product, err0 := repo.q.GetProductDetailByID(ctx.Context(), pgProductUUID)
	if err0 != nil {
		log.Debugf("Error Getting Product: %v", err0)
		return messages.InternalServerError("Error Getting Product")
	}

	priceAgg := product.Price * float64(quantity)

	err0 = repo.q.InsertIntoOrdersTable(ctx.Context(), db.InsertIntoOrdersTableParams{
		ID:         pgOrderUUID,
		UserID:     userId,
		Status:     db.OrderStatusEnum("initial"),
		TotalPrice: priceAgg,
		PaymentID:  paymentUUID,
	})
	if err0 != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err0)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}

	err1 := repo.q.InsertIntoOrderItemsTable(ctx.Context(), db.InsertIntoOrderItemsTableParams{
		UserID:          userId,
		ProductID:       pgProductUUID,
		ProductQuantity: quantity,
		ProductPriceAgg: priceAgg,
		OrderID:         pgOrderUUID,
	})
	if err1 != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err1)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}
	return nil
}

func (repo *UserRepository) UpdateOrderPaymentId(ctx *fiber.Ctx, orderId, paymentID uuid.UUID) *messages.AppError {

	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)
	pgPaymentUUID := utils.ConvertUUIDToPgType(paymentID)
	repo.q.UpdateOrderStatusByID(ctx.Context(), db.UpdateOrderStatusByIDParams{
		ID:     pgOrderUUID,
		Status: db.OrderStatusEnum("paid"),
	})
	err := repo.q.UpdateOrderPaymentId(ctx.Context(), db.UpdateOrderPaymentIdParams{
		ID:        pgOrderUUID,
		PaymentID: pgPaymentUUID,
	})
	if err != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}
	return nil
}
