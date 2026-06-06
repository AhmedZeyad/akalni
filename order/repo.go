package order

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type OrderRepo interface {
	GetOrderById(id int64) (Order, error)
	UpdateOrderStatus(id int64, status string) error
	GetOrderDetailsByOrderId(id int64) ([]OrdersDetails, error)
}

type orderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) OrderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) GetOrderById(id int64) (Order, error) {
	order := Order{}
	err := r.db.Get(&order, `
	SELECT
		id,
		client_id,
		rest_id,
		total_price,
		sub_total_price,
		last_status
	FROM orders
	WHERE id = $1
	limit 1`, id)
	if err != nil {
		slog.Error("failed to get order by id", "id", id, "error", err)
		return Order{}, err
	}
	return order, nil
}

func (r *orderRepo) GetOrderDetailsByOrderId(id int64) ([]OrdersDetails, error) {
	orderDetails := []OrdersDetails{}
	err := r.db.Select(&orderDetails, `
		SELECT
			id,
			order_id,
			product_id,
			quantity,
			price
		FROM order_details
		WHERE order_id = $1
		order by id DESC
		`, id)
	if err != nil {
		slog.Error("failed to get order details by order id", "id", id, "error", err)
		return nil, err
	}
	return orderDetails, nil
}

// TODO update update by from token id
func (r *orderRepo) UpdateOrderStatus(id int64, status string) error {
	_, err := r.db.Exec(`
		update orders
			set last_status = $2
		where id = $1
		and last_status != $2
		`, id, status)
	if err != nil {
		slog.Error("failed to update order status", "id", id, "status", status, "error", err)
		return err
	}
	return nil
}
