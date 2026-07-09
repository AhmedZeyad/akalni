package categories

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/gin-gonic/gin"
)

type categoriesHandler struct {
	categoriesService
}

func NewCategoriesHandler(service categoriesService) *categoriesHandler {
	return &categoriesHandler{categoriesService: service}
}
func (h *categoriesHandler) CreateCategory(ctx *gin.Context) {
	var (
		req    CategoriesRequest
		appErr shared.AppError
	)
	if err := ctx.BindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind body", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}

	createdBy := ctx.MustGet("user").(middleware.AdminClaims).ID

	if err := h.categoriesService.Create(int(createdBy), req); err != nil {
		appErr.Error = err
		slog.Error("failed to create category", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}

	shared.Respond(ctx, nil, nil)
}

func (h *categoriesHandler) GetCategories(ctx *gin.Context) {
	var (
		appErr shared.AppError
		req    CategoriesSearchParam
	)
	if err := ctx.BindQuery(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind query", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}

	categories, count, err := h.categoriesService.Search(req)

	if err != nil {
		appErr.Error = err
		slog.Error("failed to search categories", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, shared.PaginationResponse{Count: count, Result: categories}, nil)
}

func (h *categoriesHandler) UpdateCategory(ctx *gin.Context) {
	var (
		req    CategoriesRequest
		appErr shared.AppError
	)
	if err := ctx.BindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind body", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}

	updatedBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	if err := h.categoriesService.Update(int(updatedBy), req); err != nil {
		appErr.Error = err
		slog.Error("failed to update category", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, nil, nil)
}

func (h *categoriesHandler) DeleteCategory(ctx *gin.Context) {
	var (
		appErr shared.AppError
		req    CategoriesRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind body", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	deletedBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	if err := h.categoriesService.Delete(req.ID, int(deletedBy)); err != nil {
		appErr.Error = err
		slog.Error("failed to delete category", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, nil, nil)
}

func (h *categoriesHandler) UpdateCategoryStatus(ctx *gin.Context) {
	var (
		req    CategoriesRequest
		appErr shared.AppError
	)
	if err := ctx.BindJSON(&req); err != nil {
		appErr.Error = err
		slog.Error("failed to bind body", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	updatedBy := ctx.MustGet("user").(middleware.AdminClaims).ID
	if err := h.categoriesService.UpdateStatus(int(updatedBy), req); err != nil {
		appErr.Error = err
		slog.Error("failed to update category status", "error", err)
		shared.Respond(ctx, nil, &appErr)
		return
	}
	shared.Respond(ctx, nil, nil)
}
