package order

import (
	"errors"
	"log/slog"
	"math"
)

type OrderService struct {
	orderRepo OrderRepo
}

func NewOrderService(orderRepo OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) GetOrderById(id int64) (Order, []OrdersDetails, error) {

	order, err := s.orderRepo.GetOrderById(id)
	if err != nil {
		slog.Error("failed to get order by id", "id", id, "error", err)
		return Order{}, nil, err
	}
	orderDetails, err := s.orderRepo.GetOrderDetailsByOrderId(order.ID)
	if err != nil {
		slog.Error("failed to get order details by order id", "order_id", order.ID, "error", err)
		return Order{}, nil, err
	}
	return order, orderDetails, nil
}
func (s *OrderService) UpdateOrderStatus(id int64, status string) error {
	err := s.orderRepo.UpdateOrderStatus(id, status)
	if err != nil {
		slog.Error("failed to update order status", "id", id, "status", status, "error", err)
		return err
	}
	return nil
}
func (os *OrderService) CreateOrder(order CreateOrderRequest, restaurant Restaurant) (int64, error) {
	if len(restaurant.Products) != len(order.Products) {
		err := errors.New("some products are not available")
		slog.Error("products and order products length mismatch", "error", err)
		return 0, err
	}
	var deliveryFee = func(price float64) float64 {
		return math.Max(price*(10/100), 1000)
	}

	for i := range restaurant.Products {
		order.Subtotal += restaurant.Products[i].Price * float64(order.Products[i].Quantity)
	}

	order.TotalPrice = order.Subtotal + deliveryFee(order.Subtotal)

	order.TotalPrice = math.Ceil(order.Subtotal/250) * 250

	id, err := os.orderRepo.Create(order)
	if err != nil {
		slog.Error("failed to create order", "error", err)
		return 0, err
	}
	return id, nil
}
