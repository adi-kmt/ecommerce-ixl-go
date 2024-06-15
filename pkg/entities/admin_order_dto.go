package entities

import (
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
)

type AdminOrderDto struct {
	OrderId     string
	OrderStatus string
}

func AdminOrderDtoFromDb(orderRows []*db.GetOrdersByUserIDOrStatusRow) []AdminOrderDto {
	var orders []AdminOrderDto
	for _, order := range orderRows {
		orders = append(orders, AdminOrderDto{
			OrderId:     utils.ConvertPgUUIDToString(order.ID),
			OrderStatus: string(order.Status),
		})
	}
	return orders
}
