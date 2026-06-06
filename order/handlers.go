package order

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *OrderService
}

func NewOrderHandler(orderService *OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) GetOrder(ctx *gin.Context) {
	appError := shared.AppError{}
	var req OrderRequest
	if err := ctx.BindQuery(&req); err != nil {
		appError.Error = err
		slog.Error("failed to bind query", "error", err)
		shared.Respond(ctx, nil, &appError)
		return
	}
	order, orderDetails, err := h.orderService.GetOrderById(req.ID)
	if err != nil {
		appError.Error = err
		slog.Error("failed to get order", "error", err)
		shared.Respond(ctx, nil, &appError)
		return
	}

	shared.Respond(ctx, toOrderResponse(order, orderDetails), nil)
}

func (h *OrderHandler) UpdateOrderStatus(ctx *gin.Context) {
	appError := shared.AppError{}
	var req OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", err)
		shared.Respond(ctx, nil, &appError)
		return
	}
	err := h.orderService.UpdateOrderStatus(req.ID, req.NewStatus)
	if err != nil {
		appError.Error = err
		slog.Error("failed to update order status", "error", err)
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, nil, nil)
}
