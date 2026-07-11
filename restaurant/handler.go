package restaurant

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	service RestaurantService
	// jwt     *jwt.Token
}

func NewRestaurantHandler(service RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{
		service: service,
		// jwt:     jwt,
	}
}

func (h *RestaurantHandler) CreateRestaurant(ctx *gin.Context) {
	var req RestaurantRequest
	appError := shared.AppError{Layer: "Handler"}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind json", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	user := ctx.MustGet("user").(middleware.AdminClaims)

	err := h.service.CreateRestaurant(req, int(user.ID))
	if err != nil {
		slog.Error("failed to create restaurant", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, nil, nil)
}
func (h *RestaurantHandler) UpdateRestaurant(ctx *gin.Context) {
	var req RestaurantRequest
	appError := shared.AppError{Layer: "Handler"}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind json", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	user := ctx.MustGet("user").(middleware.AdminClaims)

	err := h.service.UpdateRestaurant(req, int(user.ID))
	if err != nil {
		slog.Error("failed to update restaurant", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, nil, nil)
}
func (h *RestaurantHandler) UpdateRestaurantStatus(ctx *gin.Context) {
	var req RestaurantRequest
	appError := shared.AppError{Layer: "Handler"}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind json", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	user := ctx.MustGet("user").(middleware.AdminClaims)

	err := h.service.UpdateRestaurantStatus(req, int(user.ID))
	if err != nil {
		slog.Error("failed to update restaurant", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, nil, nil)
}
func (h *RestaurantHandler) SearchRestaurant(ctx *gin.Context) {
	var req SearchRequest
	appError := shared.AppError{Layer: "Handler"}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind json", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}

	res, err := h.service.SearchRestaurant(req)
	if err != nil {
		slog.Error("failed to update restaurant", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, res, nil)
}

func (h *RestaurantHandler) GetRestaurantById(ctx *gin.Context) {
	var req SearchRequest
	appError := shared.AppError{Layer: "Handler"}
	if err := ctx.BindQuery(&req); err != nil {
		slog.Error("failed to bind query", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}

	res, err := h.service.SearchRestaurant(req)
	if err != nil {
		slog.Error("failed to get restaurant by id", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, res, nil)
}

func (h *RestaurantHandler) GetActiveRestaurant(ctx *gin.Context) {
	var req SearchRequest
	appError := shared.AppError{Layer: "Handler"}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		slog.Error("failed to bind query", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	res, err := h.service.GetActiveRestaurant(req)
	if err != nil {
		slog.Error("failed to get active restaurant", "error", err)
		appError.Error = err
		shared.Respond(ctx, nil, &appError)
		return
	}
	shared.Respond(ctx, res, nil)
}
