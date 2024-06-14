// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: all_sql.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const changeProductDetailsByID = `-- name: ChangeProductDetailsByID :exec
UPDATE products 
SET 
    name = COALESCE($2, name),
    description = COALESCE($3, description),
    price = COALESCE($4, price),
    stock = COALESCE($5, stock),
    category_id = COALESCE($6, category_id)
WHERE id = $1
`

type ChangeProductDetailsByIDParams struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Stock       int16       `json:"stock"`
	CategoryID  int16       `json:"category_id"`
}

func (q *Queries) ChangeProductDetailsByID(ctx context.Context, arg ChangeProductDetailsByIDParams) error {
	_, err := q.db.Exec(ctx, changeProductDetailsByID,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.CategoryID,
	)
	return err
}

const deleteProductByID = `-- name: DeleteProductByID :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProductByID(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProductByID, id)
	return err
}

const getCurrentOrderByID = `-- name: GetCurrentOrderByID :one
SELECT id, user_id, status, payment_id, total_price FROM orders
WHERE id = $1
`

func (q *Queries) GetCurrentOrderByID(ctx context.Context, id pgtype.UUID) (*Order, error) {
	row := q.db.QueryRow(ctx, getCurrentOrderByID, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.PaymentID,
		&i.TotalPrice,
	)
	return &i, err
}

const getOrdersByUserIDOrStatus = `-- name: GetOrdersByUserIDOrStatus :many
SELECT id, user_id, status, payment_id, total_price FROM orders
WHERE ($1 IS NULL OR user_id = $1)
AND ($2 IS NULL OR status = $2)
`

type GetOrdersByUserIDOrStatusParams struct {
	Column1 interface{} `json:"column_1"`
	Column2 interface{} `json:"column_2"`
}

func (q *Queries) GetOrdersByUserIDOrStatus(ctx context.Context, arg GetOrdersByUserIDOrStatusParams) ([]*Order, error) {
	rows, err := q.db.Query(ctx, getOrdersByUserIDOrStatus, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Status,
			&i.PaymentID,
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

const getProductDetailByID = `-- name: GetProductDetailByID :one
SELECT id, name, description, price, stock, category_id FROM products
WHERE id = $1
`

func (q *Queries) GetProductDetailByID(ctx context.Context, id pgtype.UUID) (*Product, error) {
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
WHERE category_id IN (SELECT id FROM categories WHERE name = ANY ($1::text[]))
`

func (q *Queries) GetProductsForCategories(ctx context.Context, dollar_1 []string) ([]*Product, error) {
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
SELECT users.id, email, name, address, isadmin, password, orders.id, user_id, status, payment_id, total_price FROM users
LEFT JOIN orders ON users.id = orders.user_id
WHERE users.id = $1
`

type GetUserDetailsAndOrdersRow struct {
	ID         int64               `json:"id"`
	Email      string              `json:"email"`
	Name       string              `json:"name"`
	Address    string              `json:"address"`
	Isadmin    bool                `json:"isadmin"`
	Password   string              `json:"password"`
	ID_2       pgtype.UUID         `json:"id_2"`
	UserID     *int64              `json:"user_id"`
	Status     NullOrderStatusEnum `json:"status"`
	PaymentID  pgtype.UUID         `json:"payment_id"`
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
			&i.Isadmin,
			&i.Password,
			&i.ID_2,
			&i.UserID,
			&i.Status,
			&i.PaymentID,
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

const insertIntoCategoriesTable = `-- name: InsertIntoCategoriesTable :exec
INSERT INTO categories (id, name)
    VALUES ($1, $2) ON CONFLICT DO NOTHING
`

type InsertIntoCategoriesTableParams struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) InsertIntoCategoriesTable(ctx context.Context, arg InsertIntoCategoriesTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoCategoriesTable, arg.ID, arg.Name)
	return err
}

const insertIntoOrderItemsTable = `-- name: InsertIntoOrderItemsTable :exec
INSERT INTO orderitems (id, user_id, product_id, product_quantity, product_price_agg, order_id)
    VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING
`

type InsertIntoOrderItemsTableParams struct {
	ID              pgtype.UUID `json:"id"`
	UserID          int64       `json:"user_id"`
	ProductID       pgtype.UUID `json:"product_id"`
	ProductQuantity int16       `json:"product_quantity"`
	ProductPriceAgg float64     `json:"product_price_agg"`
	OrderID         pgtype.UUID `json:"order_id"`
}

func (q *Queries) InsertIntoOrderItemsTable(ctx context.Context, arg InsertIntoOrderItemsTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoOrderItemsTable,
		arg.ID,
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
INSERT INTO products (id, name, description, price, stock, category_id)
    VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING
`

type InsertIntoProductsTableParams struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Stock       int16       `json:"stock"`
	CategoryID  int16       `json:"category_id"`
}

func (q *Queries) InsertIntoProductsTable(ctx context.Context, arg InsertIntoProductsTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoProductsTable,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.CategoryID,
	)
	return err
}

const insertIntoUsersTable = `-- name: InsertIntoUsersTable :exec
INSERT INTO users (id, email, name, address, isAdmin, password)
    VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING
`

type InsertIntoUsersTableParams struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Isadmin  bool   `json:"isadmin"`
	Password string `json:"password"`
}

func (q *Queries) InsertIntoUsersTable(ctx context.Context, arg InsertIntoUsersTableParams) error {
	_, err := q.db.Exec(ctx, insertIntoUsersTable,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.Address,
		arg.Isadmin,
		arg.Password,
	)
	return err
}

const searchProducts = `-- name: SearchProducts :many
SELECT id, name, description, price, stock, category_id FROM products
WHERE name ILIKE '%' || $1 || '%'
AND ($2::text[] IS NULL OR category_id IN (SELECT id FROM categories WHERE name = ANY ($2)))
`

type SearchProductsParams struct {
	Column1 *string  `json:"column_1"`
	Column2 []string `json:"column_2"`
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