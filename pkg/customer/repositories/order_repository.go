package user_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	tx, err := repo.conn.BeginTx(ctx.Context(), pgx.TxOptions{})

	if err != nil {
		log.Debugf("Error Creating Txn: %v", err)
		return messages.InternalServerError("Error Connecting to DB")
	}

	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)

	product, err0 := repo.q.WithTx(tx).GetProductDetailByID(ctx.Context(), productId)
	if err0 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Getting Product: %v", err0)
		return messages.InternalServerError("Error Getting Product")
	}

	priceAgg := product.Price * float64(quantity)
	if product.Stock < quantity {
		tx.Rollback(ctx.Context())
		return messages.InternalServerError("Out of Stock")
	}
	err1 := repo.q.WithTx(tx).InsertIntoOrderItemsTable(ctx.Context(), db.InsertIntoOrderItemsTableParams{
		OrderID:         pgOrderUUID,
		ProductID:       productId,
		UserID:          userId,
		ProductQuantity: quantity,
		ProductPriceAgg: priceAgg,
	})
	if err1 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Inserting Item Into Order: %v", err1)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}

	err2 := repo.q.WithTx(tx).UpdateProductStock(ctx.Context(), db.UpdateProductStockParams{
		ID:    productId,
		Stock: product.Stock - quantity,
	})
	if err2 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Updating Product Stock: %v", err2)
		return messages.InternalServerError("Error Updating Product Stock")
	}

	orderData, err3 := repo.q.WithTx(tx).GetOrderDetailsById(ctx.Context(), pgOrderUUID)
	if err3 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Getting Order Data: %v", err3)
		return messages.InternalServerError("Error Getting Order Data")
	}
	err4 := repo.q.WithTx(tx).UpdateOrderTotalPriceByID(ctx.Context(), db.UpdateOrderTotalPriceByIDParams{
		ID:         pgOrderUUID,
		TotalPrice: orderData.TotalPrice + priceAgg,
	})
	if err4 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Updating Order Total Price: %v", err4)
		return messages.InternalServerError("Error Updating Order Total Price")
	}
	err5 := tx.Commit(ctx.Context())
	if err5 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Committing Txn: %v", err5)
		return messages.InternalServerError("Error Committing Txn")
	}
	return nil
}

func (repo *UserRepository) InsertIntoOrderAndOrderItems(ctx *fiber.Ctx, productId int64, userId int64, quantity int16) (string, *messages.AppError) {
	tx, err5 := repo.conn.BeginTx(ctx.Context(), pgx.TxOptions{})
	if err5 != nil {
		log.Debugf("Error Creating Txn: %v", err5)
		return "", messages.InternalServerError("Error Connecting to DB")
	}

	orderId, err := utils.GenerateNewUUID()
	if err != nil {
		tx.Rollback(ctx.Context())
		return "", messages.InternalServerError("Error Generating UUID")
	}
	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)
	paymentUUID := utils.ConvertUUIDToPgType(uuid.Nil)
	product, err0 := repo.q.WithTx(tx).GetProductDetailByID(ctx.Context(), productId)
	if err0 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Getting Product: %v", err0)
		return "", messages.InternalServerError("Error Getting Product")
	}
	if product.Stock < quantity {
		tx.Rollback(ctx.Context())
		return "", messages.InternalServerError("Out of Stock")
	}

	priceAgg := product.Price * float64(quantity)

	err0 = repo.q.WithTx(tx).InsertIntoOrdersTable(ctx.Context(), db.InsertIntoOrdersTableParams{
		ID:         pgOrderUUID,
		UserID:     userId,
		Status:     db.OrderStatusEnumINITIAL,
		TotalPrice: priceAgg,
		PaymentID:  paymentUUID,
	})
	if err0 != nil {
		log.Debugf("Error Inserting Item Into Order: %v", err0)
		return "", messages.InternalServerError("Error Inserting Item Into Order")
	}

	err1 := repo.q.WithTx(tx).InsertIntoOrderItemsTable(ctx.Context(), db.InsertIntoOrderItemsTableParams{
		UserID:          userId,
		ProductID:       productId,
		ProductQuantity: quantity,
		ProductPriceAgg: priceAgg,
		OrderID:         pgOrderUUID,
	})
	if err1 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Inserting Item Into Order: %v", err1)
		return "", messages.InternalServerError("Error Inserting Item Into Order")
	}

	err2 := repo.q.WithTx(tx).UpdateProductStock(ctx.Context(), db.UpdateProductStockParams{
		ID:    productId,
		Stock: product.Stock - quantity,
	})
	if err2 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Updating Product Stock: %v", err2)
		return "", messages.InternalServerError("Error Updating Product Stock")
	}
	err3 := tx.Commit(ctx.Context())
	if err3 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Committing Txn: %v", err3)
		return "", messages.InternalServerError("Error Committing Txn")
	}
	return utils.ConvertPgUUIDToString(pgOrderUUID), nil
}

func (repo *UserRepository) UpdateOrderPaymentId(ctx *fiber.Ctx, orderId, paymentID uuid.UUID) *messages.AppError {
	tx, err2 := repo.conn.BeginTx(ctx.Context(), pgx.TxOptions{})
	if err2 != nil {
		log.Debugf("Error Creating Txn: %v", err2)
		return messages.InternalServerError("Error Connecting to DB")
	}

	pgOrderUUID := utils.ConvertUUIDToPgType(orderId)
	pgPaymentUUID := utils.ConvertUUIDToPgType(paymentID)
	err0 := repo.q.WithTx(tx).UpdateOrderStatusByID(ctx.Context(), db.UpdateOrderStatusByIDParams{
		ID:     pgOrderUUID,
		Status: db.OrderStatusEnumPAID,
	})
	if err0 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Inserting Item Into Order: %v", err0)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}
	err := repo.q.WithTx(tx).UpdateOrderPaymentId(ctx.Context(), db.UpdateOrderPaymentIdParams{
		ID:        pgOrderUUID,
		PaymentID: pgPaymentUUID,
	})
	if err != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Inserting Item Into Order: %v", err)
		return messages.InternalServerError("Error Inserting Item Into Order")
	}
	err3 := tx.Commit(ctx.Context())
	if err3 != nil {
		tx.Rollback(ctx.Context())
		log.Debugf("Error Committing Txn: %v", err3)
		return messages.InternalServerError("Error Committing Txn")
	}
	return nil
}
