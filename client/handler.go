package client

import (
	"log/slog"
	"net/http"

	"github.com/AhmedZeyad/Akalni/auth"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service ClientService
}

func NewClientHandler(service ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h ClientHandler) GetProfile(ctx *gin.Context) {
	value, ok := ctx.Get("client")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		slog.Error("error on get profile", "error", "client not found", "value", value)
		return
	}
	slog.Error("error on get profile", "error", "client not found", "value", value)
	claims := value.(auth.Claims)

	client, err := h.service.GetProfile(claims.ClientID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error("error on get profile", "error", err)
		return
	}
	ctx.JSON(http.StatusOK, client)
}
