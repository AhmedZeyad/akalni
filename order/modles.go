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

type ProductRequest struct {
	ID       int `json:"id"`
	Quantity int `json:"qty"`
}
type ProductData struct {
	ID    int64
	Price float64
}
type RestaurantData struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type CreateOrderRequest struct {
	ClientID     int64            `json:"-"`
	RestaurantID int              `json:"rest_id"`
	Products     []ProductRequest `json:"products"`
	TotalPrice   float64          `json:"-"`
	Subtotal     float64          `json:"-"`
}

type OrderDataProvider interface {
	GetRestaurantByID(restID int, productsIDS []int) (Restaurant, error)
}
type Restaurant struct {
	ID       int64
	Name     string
	Status   bool
	Lon      float64
	Lat      float64
	Address  string
	Products []ProductData
}
