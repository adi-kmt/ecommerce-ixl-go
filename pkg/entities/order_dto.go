package entities

import (
	"github.com/google/uuid"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/utils"
)

type orderLineItems struct {
	ProductId string  `json:"product_id"`
	Quantity  int16   `json:"quantity"`
	PriceAgg  float64 `json:"price_agg"`
}

type OrderDto struct {
	OrderId        string           `json:"order_id"`
	OrderLineItems []orderLineItems `json:"order_line_items"`
}

func OrderDtoFromOrderDb(order []*db.GetCurrentOrderByIDRow, orderId uuid.UUID) *OrderDto {
	var orderItems []orderLineItems

	for _, item := range order {
		orderItems = append(orderItems, orderLineItems{
			ProductId: utils.ConvertPgUUIDToString(item.ProductID),
			Quantity:  item.ProductQuantity,
			PriceAgg:  item.ProductPriceAgg,
		})
	}
	return &OrderDto{
		OrderId:        orderId.String(),
		OrderLineItems: orderItems,
	}
}
