// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type OrderStatusEnum string

const (
	OrderStatusEnumPENDING   OrderStatusEnum = "PENDING"
	OrderStatusEnumONTHEWAY  OrderStatusEnum = "ON THE WAY"
	OrderStatusEnumCOMPLETED OrderStatusEnum = "COMPLETED"
	OrderStatusEnumCANCELLED OrderStatusEnum = "CANCELLED"
)

func (e *OrderStatusEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatusEnum(s)
	case string:
		*e = OrderStatusEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatusEnum: %T", src)
	}
	return nil
}

type NullOrderStatusEnum struct {
	OrderStatusEnum OrderStatusEnum `json:"order_status_enum"`
	Valid           bool            `json:"valid"` // Valid is true if OrderStatusEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatusEnum) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatusEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatusEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatusEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatusEnum), nil
}

type Category struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID         pgtype.UUID     `json:"id"`
	UserID     int64           `json:"user_id"`
	Status     OrderStatusEnum `json:"status"`
	PaymentID  pgtype.UUID     `json:"payment_id"`
	TotalPrice float64         `json:"total_price"`
}

type Orderitem struct {
	ID              pgtype.UUID `json:"id"`
	UserID          int64       `json:"user_id"`
	ProductID       pgtype.UUID `json:"product_id"`
	ProductQuantity int16       `json:"product_quantity"`
	ProductPriceAgg float64     `json:"product_price_agg"`
	OrderID         pgtype.UUID `json:"order_id"`
}

type Product struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Stock       int16       `json:"stock"`
	CategoryID  int16       `json:"category_id"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Isadmin  bool   `json:"isadmin"`
	Password string `json:"password"`
}
