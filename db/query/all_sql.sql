-- name: InsertIntoProductsTable :exec
INSERT INTO products (id, name, description, price, stock, category_id)
    VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING;

-- name: InsertIntoCategoriesTable :exec
INSERT INTO categories (id, name)
    VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: InsertIntoUsersTable :exec
INSERT INTO users (id, email, name, address, isAdmin, password)
    VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING;

-- name: InsertIntoOrdersTable :exec
INSERT INTO orders (id, user_id, status, payment_id, total_price)
    VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING;

-- name: InsertIntoOrderItemsTable :exec
INSERT INTO orderitems (id, user_id, product_id, product_quantity, product_price_agg, order_id)
    VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING;

-- name: SearchProducts :many
SELECT * FROM products
WHERE name ILIKE '%' || $1 || '%'
AND ($2::text[] IS NULL OR category_id IN (SELECT id FROM categories WHERE name = ANY ($2)));

-- name: GetUserDetailsAndOrders :many
SELECT * FROM users
LEFT JOIN orders ON users.id = orders.user_id
WHERE users.id = $1;

-- name: GetCurrentOrderByID :one
SELECT * FROM orders
WHERE id = $1;

-- name: GetProductsForCategories :many
SELECT * FROM products
WHERE category_id IN (SELECT id FROM categories WHERE name = ANY ($1::text[]));

-- name: GetProductDetailByID :one
SELECT * FROM products
WHERE id = $1;

-- name: GetOrdersByUserIDOrStatus :many
SELECT * FROM orders
WHERE ($1 IS NULL OR user_id = $1)
AND ($2 IS NULL OR status = $2);

-- name: UpdateOrderStatusByID :exec
UPDATE orders SET status = $2 WHERE id = $1;

-- name: DeleteProductByID :exec
DELETE FROM products WHERE id = $1;

-- name: ChangeProductDetailsByID :exec
UPDATE products 
SET 
    name = COALESCE($2, name),
    description = COALESCE($3, description),
    price = COALESCE($4, price),
    stock = COALESCE($5, stock),
    category_id = COALESCE($6, category_id)
WHERE id = $1;