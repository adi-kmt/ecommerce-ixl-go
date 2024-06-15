package user_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

func (repo *UserRepository) GetItemsInCart(ctx *fiber.Ctx, orderId uuid.UUID) (*entities.OrderDto, *messages.AppError) {

	pgUUID := utils.ConvertUUIDToPgType(orderId)
	order, err := repo.q.GetCurrentOrderByID(ctx.Context(), pgUUID)
	if err != nil {
		log.Debugf("Error Getting Items In Cart: %v", err)
		return nil, messages.InternalServerError("Error Getting Items In Cart")
	}
	return entities.OrderDtoFromOrderDb(order, orderId), nil
}

func (repo *UserRepository) InsertItemIntoOrderItem(ctx *fiber.Ctx, productId int64, orderId uuid.UUID, userId int64, quantity int16) *messages.AppError {

	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)

	product, err0 := repo.q.GetProductDetailByID(ctx.Context(), productId)
	if err0 != nil {
		log.Debugf("Error Getting Product: %v", err0)
		return messages.InternalServerError("Error Getting Product")
	}

	priceAgg := product.Price * float64(quantity)
	if product.Stock < quantity {
		return messages.InternalServerError("Out of Stock")
	}
	err1 := repo.q.InsertIntoOrderItemsTable(ctx.Context(), db.InsertIntoOrderItemsTableParams{
		OrderID:         pgOrderUUID,
		ProductID:       productId,
		UserID:          userId,
		ProductQuantity: quantity,
		ProductPriceAgg: priceAgg,
	})
	if err1 != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err1)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}

	err2 := repo.q.UpdateProductStock(ctx.Context(), db.UpdateProductStockParams{
		ID:    productId,
		Stock: product.Stock - quantity,
	})
	if err2 != nil {
		log.Debugf("Error Updating Product Stock: %v", err2)
		return messages.InternalServerError("Error Updating Product Stock")
	}
	return nil
}

func (repo *UserRepository) InsertIntoOrderAndOrderItems(ctx *fiber.Ctx, productId int64, userId int64, quantity int16) (string, *messages.AppError) {
	orderId, err := utils.GenerateNewUUID()
	if err != nil {
		return "", messages.InternalServerError("Error Generating UUID")
	}
	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)
	paymentUUID := utils.ConvertUUIDToPgType(uuid.Nil)
	product, err0 := repo.q.GetProductDetailByID(ctx.Context(), productId)
	if err0 != nil {
		log.Debugf("Error Getting Product: %v", err0)
		return "", messages.InternalServerError("Error Getting Product")
	}
	if product.Stock < quantity {
		return "", messages.InternalServerError("Out of Stock")
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
		return "", messages.InternalServerError("Error Inserting Item Into Order")
	}

	err1 := repo.q.InsertIntoOrderItemsTable(ctx.Context(), db.InsertIntoOrderItemsTableParams{
		UserID:          userId,
		ProductID:       productId,
		ProductQuantity: quantity,
		ProductPriceAgg: priceAgg,
		OrderID:         pgOrderUUID,
	})
	if err1 != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err1)
		return "", messages.InternalServerError("Error Inserting Item Into Order")
	}

	err2 := repo.q.UpdateProductStock(ctx.Context(), db.UpdateProductStockParams{
		ID:    productId,
		Stock: product.Stock - quantity,
	})
	if err2 != nil {
		log.Debugf("Error Updating Product Stock: %v", err2)
		return "", messages.InternalServerError("Error Updating Product Stock")
	}
	return utils.ConvertPgUUIDToString(pgOrderUUID), nil
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
