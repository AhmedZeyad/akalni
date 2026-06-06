package order

type Order struct {
	ID           int64   `db:"id"`
	ClientID     int64   `db:"client_id"`
	RestaurantID int64   `db:"rest_id"`
	TotalPrice   float64 `db:"total_price"`
	Subtotal     float64 `db:"sub_total_price"`
	LastStatus   string  `db:"last_status"`
}

type OrdersDetails struct {
	ID        int64   `db:"id"`
	OrderID   int64   `db:"order_id"`
	ProductID int64   `db:"product_id"`
	Quantity  int64   `db:"quantity"`
	Price     float64 `db:"price"`
}
type orderItem struct {
	ID       int64   `db:"id"`
	Quantity int64   `db:"quantity"`
	Price    float64 `db:"price"`
}
type OrderResponse struct {
	Order `json:"order"`
	Items []orderItem `json:"items"`
}

func toOrderResponse(order Order, items []OrdersDetails) OrderResponse {
	orderItems := make([]orderItem, len(items))
	for i, item := range items {
		orderItems[i] = orderItem{
			ID:       item.ID,
			Quantity: item.Quantity,
			Price:    item.Price,
		}
	}
	return OrderResponse{
		Order: order,
		Items: orderItems,
	}
}

type OrderRequest struct {
	ID        int64  `form:"id" json:"id"`
	NewStatus string `json:"new_status"`
}
