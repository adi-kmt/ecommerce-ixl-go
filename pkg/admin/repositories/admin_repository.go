package admin_repositories

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

type AdminRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewAdminRepository(conn *pgxpool.Pool, q *db.Queries) *AdminRepository {
	return &AdminRepository{
		q:    q,
		conn: conn,
	}
}

func (repo *AdminRepository) AddProduct(ctx *fiber.Ctx, name string, price float64, categoryID int32, stock int16) *messages.AppError {
	err1 := repo.q.InsertIntoProductsTable(ctx.Context(), db.InsertIntoProductsTableParams{
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
	err := repo.q.InsertIntoCategoriesTable(ctx.Context(), name)
	if err != nil {
		log.Debugf("Error Inserting Category: %v", err)
		return messages.InternalServerError("Error Inserting Category")
	}
	return nil
}

func (repo *AdminRepository) DeleteProduct(ctx *fiber.Ctx, id int64) *messages.AppError {
	err := repo.q.DeleteProductByID(ctx.Context(), id)
	if err != nil {
		log.Debugf("Error Deleting Product: %v", err)
		return messages.InternalServerError("Error Deleting Product")
	}
	return nil
}

func (repo *AdminRepository) GetAllOrders(ctx *fiber.Ctx, userId string, status string) ([]entities.AdminOrderDto, *messages.AppError) {
	var userIntVal int32 = -1
	var pgStatus db.OrderStatusEnum = ""

	if userId != "" {
		intval, err := strconv.Atoi(userId)
		if err != nil {
			log.Debugf("Error Parsing User ID: %v", err)
			return nil, messages.BadRequest("Invalid User ID")
		}
		int32Val := int32(intval)
		userIntVal = int32Val
	}

	if status != "" {
		pgStatusVal := db.OrderStatusEnum(status)
		pgStatus = pgStatusVal
	}

	orders, err0 := repo.q.GetOrdersByUserIDOrStatus(ctx.Context(), db.GetOrdersByUserIDOrStatusParams{
		Column1: userIntVal,
		Column2: pgStatus,
	})
	if err0 != nil {
		log.Debugf("Error Getting Orders: %v", err0)
		return nil, messages.InternalServerError("Error Getting Orders")
	}

	return entities.AdminOrderDtoFromDb(orders), nil
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
