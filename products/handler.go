package products

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	*productsService
}

func NewProductHandler(ps *productsService) *productHandler {
	return &productHandler{productsService: ps}
}

func (h *productHandler) GetProduct(ctx *gin.Context) {
	var (
		req    ProductsSearchParam
		appErr shared.AppError
	)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind json", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	product, count, err := h.productsService.Search(req)
	if err != nil {
		appErr.Error = err
		slog.Error("failed to get product", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, shared.PaginationResponse{Result: product, Count: count}, nil)
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	var (
		req    ProductsRequest
		appErr shared.AppError
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind json", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	createdBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	err := h.productsService.Create(int(createdBy), req)
	if err != nil {
		appErr.Error = err
		slog.Error("failed to add product", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, createdBy, nil)
}

func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var (
		req    ProductsRequest
		appErr shared.AppError
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind json", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	updatedBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	err := h.productsService.Update(int(updatedBy), req)
	if err != nil {
		appErr.Error = err
		slog.Error("failed to update product", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, nil, nil)
}

func (h *productHandler) UpdateProductStatus(ctx *gin.Context) {
	var (
		req    ProductsRequest
		appErr shared.AppError
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind json", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	updatedBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	slog.Error("updating product status", "id", req.ID, "active", req.Active)

	err := h.productsService.UpdateStatus(int(updatedBy), req)
	if err != nil {
		appErr.Error = err
		slog.Error("failed to update product status", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, nil, nil)
}

func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	var (
		req    ProductsRequest
		appErr shared.AppError
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind json", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	deletedBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	err := h.productsService.Delete(req.ID, int(deletedBy))
	if err != nil {
		appErr.Error = err
		slog.Error("failed to delete product", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, nil, nil)
}
