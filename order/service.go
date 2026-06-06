package order

import (
	"log/slog"
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
