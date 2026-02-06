package auth

import (
	"context"

	"github.com/ba7rIbrahim/Akalni/logger"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service ClientService
}

func NewClientHandler(service ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) Create(c *gin.Context) {
	logger.Log.Debug("Request info ", "path", c.Request.URL.Path, "method", c.Request.Method,
		"content-type", c.GetHeader("Content-Type"),
		"content-length", c.Request.ContentLength)
	var request RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Log.Error("binding erro", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// TODO: Implement user registration logic
	res, err := h.service.CreateUser(context.Background(), &request)
	if err != nil {
		logger.Log.Error("create user erro", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully", "data": res})
}
func (h *ClientHandler) Login(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Log.Error("binding erro", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client, err := h.service.Login(context.Background(), &request)
	if err != nil {
		logger.Log.Error("login erro", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User logged in successfully", "data": client})
}
