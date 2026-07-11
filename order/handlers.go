package order

import (
	"errors"
	"log/slog"

	"github.com/AhmedZeyad/Akalni/customErrors"
	"github.com/AhmedZeyad/Akalni/middleware"
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

func (h *OrderHandler) CreateOrder(orderDataProvider OrderDataProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			appError = shared.AppError{}
			req      CreateOrderRequest
		)

		if err := ctx.BindJSON(&req); err != nil {
			appError.Error = err
			slog.Error("failed to bind json", "error", err)
			shared.Respond(ctx, nil, &appError)
			return
		}
		if len(req.Products) == 0 {
			appError.Error = errors.New("no products provided")
			slog.Error("no products provided", "error", appError)
			shared.Respond(ctx, nil, &appError)
			return
		}
		if req.RestaurantID == 0 {
			appError.Error = errors.New("restaurant id is required")
			slog.Error("restaurant id is required", "error", appError)
			shared.Respond(ctx, nil, &appError)
			return
		}
		productsIDS := make([]int, len(req.Products))
		for i, product := range req.Products {
			productsIDS[i] = product.ID
		}
		restaurant, err := orderDataProvider.GetRestaurantByID(req.RestaurantID, productsIDS)
		if err != nil {
			appError.Error = err
			slog.Error("failed to get products", "error", err)
			shared.Respond(ctx, nil, &appError)
			return
		}
		if err != nil {
			appError.Error = err
			slog.Error("failed to get restaurant", "error", err)
			shared.Respond(ctx, nil, &appError)
			return
		}
		if len(restaurant.Products) == 0 {
			appError.Error = errors.New(customErrors.VALIDATION_REQUIRED_FIELD + " : products")
			slog.Error("no products provided", "error", appError)
			shared.Respond(ctx, nil, &appError)
			return
		}
		if restaurant.ID == 0 {
			appError.Error = errors.New(customErrors.VALIDATION_REQUIRED_FIELD + " : restaurant")
			slog.Error("restaurant id is required", "error", appError)
			shared.Respond(ctx, nil, &appError)
			return
		}
		req.ClientID = ctx.MustGet("client").(middleware.ClientClaims).ID
		id, err := h.orderService.CreateOrder(req, restaurant)
		if err != nil {
			appError.Error = err
			slog.Error("failed to create order", "error", err)
			shared.Respond(ctx, nil, &appError)
			return
		}
		shared.Respond(ctx, map[string]int64{"id": id}, nil)
	}
}
