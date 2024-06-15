package entities

import db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"

type UserOrderList struct {
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}

type UserDetailsAndOrdersDto struct {
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	OrderDetails []UserOrderList `json:"order_details"`
}

func UserDetailsAndOrdersFromDb(userRow []*db.GetUserDetailsAndOrdersRow) *UserDetailsAndOrdersDto {
	var orderDetails []UserOrderList = []UserOrderList{}
	for _, user := range userRow {
		if user.TotalPrice != nil {
			orderDetails = append(orderDetails, UserOrderList{
				Status:     string(user.Status.OrderStatusEnum),
				TotalPrice: *user.TotalPrice,
			})
		}
	}

	return &UserDetailsAndOrdersDto{
		Name:         userRow[0].Name,
		Email:        userRow[0].Email,
		OrderDetails: orderDetails,
	}
}
