package handlers

import (
	"context"

	"github.com/ba7rIbrahim/Akalni/Models"
	"github.com/ba7rIbrahim/Akalni/services"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service services.ClientService
}

func NewClientHandler(service services.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) Create(c *gin.Context) {
	var request Models.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// TODO: Implement user registration logic
	res, err := h.service.CreateUser(context.Background(), &request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully", "data": res})
}
