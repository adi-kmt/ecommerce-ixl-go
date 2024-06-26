// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: all_sql.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteProductByID = `-- name: DeleteProductByID :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProductByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProductByID, id)
	return err
}

const getCurrentOrderByID = `-- name: GetCurrentOrderByID :many
SELECT id, product_id, product_quantity, product_price_agg FROM orderitems
WHERE order_id = $1
`

type GetCurrentOrderByIDRow struct {
	ID              int32   `json:"id"`
	ProductID       int64   `json:"product_id"`
	ProductQuantity int16   `json:"product_quantity"`
	ProductPriceAgg float64 `json:"product_price_agg"`
}

func (q *Queries) GetCurrentOrderByID(ctx context.Context, orderID pgtype.UUID) ([]*GetCurrentOrderByIDRow, error) {
	rows, err := q.db.Query(ctx, getCurrentOrderByID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCurrentOrderByIDRow
	for rows.Next() {
		var i GetCurrentOrderByIDRow
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.ProductQuantity,
			&i.ProductPriceAgg,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderDetailsById = `-- name: GetOrderDetailsById :one
SELECT id, total_price FROM orders
WHERE id = $1
`

type GetOrderDetailsByIdRow struct {
	ID         pgtype.UUID `json:"id"`
	TotalPrice float64     `json:"total_price"`
}

func (q *Queries) GetOrderDetailsById(ctx context.Context, id pgtype.UUID) (*GetOrderDetailsByIdRow, error) {
	row := q.db.QueryRow(ctx, getOrderDetailsById, id)
	var i GetOrderDetailsByIdRow
	err := row.Scan(&i.ID, &i.TotalPrice)
	return &i, err
}

const getOrdersByUserIDOrStatus = `-- name: GetOrdersByUserIDOrStatus :many
SELECT id, user_id, status FROM orders
WHERE (NULLIF($1::int, -1) IS NULL OR user_id = $1::int)
AND (NULLIF($2::text, '') IS NULL OR status = $2::order_status_enum)
`

type GetOrdersByUserIDOrStatusParams struct {
	Column1 int32  `json:"column_1"`
	Column2 interface{} `json:"column_2"`
}

type GetOrdersByUserIDOrStatusRow struct {
	ID     pgtype.UUID     `json:"id"`
	UserID int64           `json:"user_id"`
	Status OrderStatusEnum `json:"status"`
}

func (q *Queries) GetOrdersByUserIDOrStatus(ctx context.Context, arg GetOrdersByUserIDOrStatusParams) ([]*GetOrdersByUserIDOrStatusRow, error) {
	rows, err := q.db.Query(ctx, getOrdersByUserIDOrStatus, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetOrdersByUserIDOrStatusRow
	for rows.Next() {
		var i GetOrdersByUserIDOrStatusRow
		if err := rows.Scan(&i.ID, &i.UserID, &i.Status); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductDetailByID = `-- name: GetProductDetailByID :one
SELECT id, name, description, price, stock, category_id FROM products
WHERE id = $1
`

func (q *Queries) GetProductDetailByID(ctx context.Context, id int64) (*Product, error) {
	row := q.db.QueryRow(ctx, getProductDetailByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.CategoryID,
	)
	return &i, err
}

const getProductsForCategories = `-- name: GetProductsForCategories :many
SELECT id, name, description, price, stock, category_id FROM products
WHERE category_id = ANY($1::int[])
`

func (q *Queries) GetProductsForCategories(ctx context.Context, dollar_1 []int32) ([]*Product, error) {
	rows, err := q.db.Query(ctx, getProductsForCategories, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Stock,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserDetailsAndOrders = `-- name: GetUserDetailsAndOrders :many
SELECT users.id, users.email, users.name, users.address, orders.status, orders.total_price FROM users
LEFT JOIN orders ON users.id = orders.user_id
WHERE users.id = $1
`

type GetUserDetailsAndOrdersRow struct {
	ID         int64               `json:"id"`
	Email      string              `json:"email"`
	Name       string              `json:"name"`
	Address    string              `json:"address"`
	Status     NullOrderStatusEnum `json:"status"`
	TotalPrice *float64            `json:"total_price"`
}

func (q *Queries) GetUserDetailsAndOrders(ctx context.Context, id int64) ([]*GetUserDetailsAndOrdersRow, error) {
	rows, err := q.db.Query(ctx, getUserDetailsAndOrders, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetUserDetailsAndOrdersRow
	for rows.Next() {
		var i GetUserDetailsAndOrdersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Name,
			&i.Address,
			&i.Status,
			&i.TotalPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserEmailAndPasswordByEmail = `-- name: GetUserEmailAndPasswordByEmail :one
SELECT  id, email, password FROM users
WHERE email = $1
`

type GetUserEmailAndPasswordByEmailRow struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) GetUserEmailAndPasswordByEmail(ctx context.Context, email string) (*GetUserEmailAndPasswordByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserEmailAndPasswordByEmail, email)
	var i GetUserEmailAndPasswordByEmailRow
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return &i, err
}

const insertIntoCategoriesTable = `-- name: InsertIntoCategoriesTable :exec
INSERT INTO categories (name)
    VALUES ($1) ON CONFLICT DO NOTHING
`

func (q *Queries) InsertIntoCategoriesTable(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, insertIntoCategoriesTable, name)
	return err
}

const insertIntoOrderItemsTable = `-- name: InsertIntoOrderItemsTable :exec
INSERT INTO orderitems (user_id, product_id, product_quantity, product_price_agg, order_id)
    VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING
`

type InsertIntoOrderItemsTableParams struct {
	UserID          int64       `json:"user_id"`
	ProductID       int64       `json:"product_id"`
	ProductQuantity int16       `json:"product_quantity"`
	ProductPriceAgg float64     `json:"product_price_agg"`
	OrderID         pgtype.UUID `json:"order_id"`
}

func (q *Queries) InsertIntoOrderItemsTable(ctx context.Context, arg InsertIntoOrderItemsTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoOrderItemsTable,
		arg.UserID,
		arg.ProductID,
		arg.ProductQuantity,
		arg.ProductPriceAgg,
		arg.OrderID,
	)
	return err
}

const insertIntoOrdersTable = `-- name: InsertIntoOrdersTable :exec
INSERT INTO orders (id, user_id, status, payment_id, total_price)
    VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING
`

type InsertIntoOrdersTableParams struct {
	ID         pgtype.UUID     `json:"id"`
	UserID     int64           `json:"user_id"`
	Status     OrderStatusEnum `json:"status"`
	PaymentID  pgtype.UUID     `json:"payment_id"`
	TotalPrice float64         `json:"total_price"`
}

func (q *Queries) InsertIntoOrdersTable(ctx context.Context, arg InsertIntoOrdersTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoOrdersTable,
		arg.ID,
		arg.UserID,
		arg.Status,
		arg.PaymentID,
		arg.TotalPrice,
	)
	return err
}

const insertIntoProductsTable = `-- name: InsertIntoProductsTable :exec
INSERT INTO products (name, description, price, stock, category_id)
    VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING
`

type InsertIntoProductsTableParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int16   `json:"stock"`
	CategoryID  int32   `json:"category_id"`
}

func (q *Queries) InsertIntoProductsTable(ctx context.Context, arg InsertIntoProductsTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoProductsTable,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.CategoryID,
	)
	return err
}

const insertIntoUsersTable = `-- name: InsertIntoUsersTable :one
INSERT INTO users (email, name, address, isAdmin, password)
    VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING RETURNING id
`

type InsertIntoUsersTableParams struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Isadmin  bool   `json:"isadmin"`
	Password string `json:"password"`
}

func (q *Queries) InsertIntoUsersTable(ctx context.Context, arg InsertIntoUsersTableParams) (int64, error) {
	row := q.db.QueryRow(ctx, insertIntoUsersTable,
		arg.Email,
		arg.Name,
		arg.Address,
		arg.Isadmin,
		arg.Password,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const searchProducts = `-- name: SearchProducts :many
SELECT id, name, description, price, stock, category_id 
FROM products
WHERE name ILIKE '%' || $1 || '%'
AND (($2::int[] IS NULL) OR (category_id = ANY($2::int[])))
`

type SearchProductsParams struct {
	Column1 *string `json:"column_1"`
	Column2 []int32 `json:"column_2"`
}

func (q *Queries) SearchProducts(ctx context.Context, arg SearchProductsParams) ([]*Product, error) {
	rows, err := q.db.Query(ctx, searchProducts, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Stock,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderPaymentId = `-- name: UpdateOrderPaymentId :exec
UPDATE orders SET payment_id = $2 WHERE id = $1
`

type UpdateOrderPaymentIdParams struct {
	ID        pgtype.UUID `json:"id"`
	PaymentID pgtype.UUID `json:"payment_id"`
}

func (q *Queries) UpdateOrderPaymentId(ctx context.Context, arg UpdateOrderPaymentIdParams) error {
	_, err := q.db.Exec(ctx, updateOrderPaymentId, arg.ID, arg.PaymentID)
	return err
}

const updateOrderStatusByID = `-- name: UpdateOrderStatusByID :exec
UPDATE orders SET status = $2 WHERE id = $1
`

type UpdateOrderStatusByIDParams struct {
	ID     pgtype.UUID     `json:"id"`
	Status OrderStatusEnum `json:"status"`
}

func (q *Queries) UpdateOrderStatusByID(ctx context.Context, arg UpdateOrderStatusByIDParams) error {
	_, err := q.db.Exec(ctx, updateOrderStatusByID, arg.ID, arg.Status)
	return err
}

const updateOrderTotalPriceByID = `-- name: UpdateOrderTotalPriceByID :exec
UPDATE orders SET total_price = $2 WHERE id = $1
`

type UpdateOrderTotalPriceByIDParams struct {
	ID         pgtype.UUID `json:"id"`
	TotalPrice float64     `json:"total_price"`
}

func (q *Queries) UpdateOrderTotalPriceByID(ctx context.Context, arg UpdateOrderTotalPriceByIDParams) error {
	_, err := q.db.Exec(ctx, updateOrderTotalPriceByID, arg.ID, arg.TotalPrice)
	return err
}

const updateProductStock = `-- name: UpdateProductStock :exec
UPDATE products SET stock = $2 WHERE id = $1
`

type UpdateProductStockParams struct {
	ID    int64 `json:"id"`
	Stock int16 `json:"stock"`
}

func (q *Queries) UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) error {
	_, err := q.db.Exec(ctx, updateProductStock, arg.ID, arg.Stock)
	return err
}
