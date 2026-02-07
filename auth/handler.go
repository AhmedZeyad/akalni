package auth

import (
	"context"

	"github.com/ba7rIbrahim/Akalni/logger"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Create(c *gin.Context) {
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
func (h *AuthHandler) Login(c *gin.Context) {
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
func (h *AuthHandler) Refresh(c *gin.Context) {
	Bearer := "Bearer "
	authHeader := c.GetHeader("Authorization")
	strToken := authHeader[len(Bearer):]
	res, err := h.service.Refresh(context.Background(), strToken)
	if err != nil {
		logger.Log.Error("refresh erro", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User refreshed successfully", "data": res})
}
