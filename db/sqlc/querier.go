// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	DeleteProductByID(ctx context.Context, id int64) error
	GetCurrentOrderByID(ctx context.Context, orderID pgtype.UUID) ([]*GetCurrentOrderByIDRow, error)
	GetOrdersByUserIDOrStatus(ctx context.Context, arg GetOrdersByUserIDOrStatusParams) ([]*GetOrdersByUserIDOrStatusRow, error)
	GetProductDetailByID(ctx context.Context, id int64) (*Product, error)
	GetProductsForCategories(ctx context.Context, dollar_1 []string) ([]*Product, error)
	GetUserDetailsAndOrders(ctx context.Context, id int64) ([]*GetUserDetailsAndOrdersRow, error)
	GetUserEmailAndPasswordByEmail(ctx context.Context, email string) (*GetUserEmailAndPasswordByEmailRow, error)
	InsertIntoCategoriesTable(ctx context.Context, name string) error
	InsertIntoOrderItemsTable(ctx context.Context, arg InsertIntoOrderItemsTableParams) error
	InsertIntoOrdersTable(ctx context.Context, arg InsertIntoOrdersTableParams) error
	InsertIntoProductsTable(ctx context.Context, arg InsertIntoProductsTableParams) error
	InsertIntoUsersTable(ctx context.Context, arg InsertIntoUsersTableParams) error
	SearchProducts(ctx context.Context, arg SearchProductsParams) ([]*Product, error)
	UpdateOrderPaymentId(ctx context.Context, arg UpdateOrderPaymentIdParams) error
	UpdateOrderStatusByID(ctx context.Context, arg UpdateOrderStatusByIDParams) error
	UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) error
}

var _ Querier = (*Queries)(nil)
