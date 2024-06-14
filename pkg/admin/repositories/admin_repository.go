package admin_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
)

type AdminRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewAdminRepository(conn *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{
		q:    db.New(conn),
		conn: conn,
	}
}

func (repo *AdminRepository) AddProduct(ctx *fiber.Ctx, name string, price float64, categoryID, stock int16) *messages.AppError {
	uuid, err := utils.GenerateNewUUID()
	if err != nil {
		return messages.InternalServerError("Error Generating UUID")
	}
	pgUuid := utils.ConvertUUIDToPgType(uuid)
	err1 := repo.q.InsertIntoProductsTable(ctx.Context(), db.InsertIntoProductsTableParams{
		ID:         pgUuid,
		Name:       name,
		Price:      price,
		Stock:      stock,
		CategoryID: categoryID,
	})
	if err1 != nil {
		log.Debugf("Error Inserting Product: %v", err1)
		return messages.InternalServerError("Error Inserting Product")
	}
	return nil
}

func (repo *AdminRepository) AddCategory(ctx *fiber.Ctx, name string) *messages.AppError {
	err := repo.q.InsertIntoCategoriesTable(ctx.Context(), db.InsertIntoCategoriesTableParams{
		Name: name,
	})
	if err != nil {
		log.Debugf("Error Inserting Category: %v", err)
		return messages.InternalServerError("Error Inserting Category")
	}
	return nil
}

func (repo *AdminRepository) DeleteProduct(ctx *fiber.Ctx, id uuid.UUID) *messages.AppError {
	pgUuid := utils.ConvertUUIDToPgType(id)
	err := repo.q.DeleteProductByID(ctx.Context(), pgUuid)
	if err != nil {
		log.Debugf("Error Deleting Product: %v", err)
		return messages.InternalServerError("Error Deleting Product")
	}
	return nil
}

func (repo *AdminRepository) GetAllOrders(ctx *fiber.Ctx, userId string, status string) *messages.AppError {
	var userUUID *uuid.UUID
	var pgStatus *db.OrderStatusEnum
	var err error

	if userId != "" {
		uuidVal, err := uuid.Parse(userId)
		if err != nil {
			log.Debugf("Error Parsing User ID: %v", err)
			return messages.BadRequest("Invalid User ID")
		}
		userUUID = &uuidVal
	}

	if status != "" {
		pgStatusVal := db.OrderStatusEnum(status)
		pgStatus = &pgStatusVal
	}

	_, err = repo.q.GetOrdersByUserIDOrStatus(ctx.Context(), db.GetOrdersByUserIDOrStatusParams{
		Column1: userUUID,
		Column2: pgStatus,
	})
	if err != nil {
		log.Debugf("Error Getting Orders: %v", err)
		return messages.InternalServerError("Error Getting Orders")
	}

	return nil
}

func (repo *AdminRepository) ChangeOrderStatus(ctx *fiber.Ctx, orderId uuid.UUID, status string) *messages.AppError {

	err := repo.q.UpdateOrderStatusByID(ctx.Context(), db.UpdateOrderStatusByIDParams{
		ID:     utils.ConvertUUIDToPgType(orderId),
		Status: db.OrderStatusEnum(status),
	})
	if err != nil {
		log.Debugf("Error Changing Order Status: %v", err)
		return messages.InternalServerError("Error Changing Order Status")
	}
	return nil
}
