package client

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/AhmedZeyad/Akalni/logger"
	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/shared"
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
	claims := value.(middleware.ClientClaims)

	client, err := h.service.GetProfile(claims.ClientID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error("error on get profile", "error", err)
		return
	}
	ctx.JSON(http.StatusOK, client)
}
func (h *ClientHandler) Create(c *gin.Context) {
	appError := shared.AppError{Layer: ""}
	logger.Log.Debug("Request info ", "path", c.Request.URL.Path, "method", c.Request.Method,
		"content-type", c.GetHeader("Content-Type"),
		"content-length", c.Request.ContentLength)
	var request RegisterRequest
	appError.Error = c.ShouldBindJSON(&request)
	if appError.Error != nil {
		logger.Log.Error("binding error", "error", appError.Error)
		shared.Respond(c, nil, &appError)
		return
	}
	// TODO: Implement user registration logic
	res, err := h.service.CreateUser(context.Background(), &request)
	if err != nil {
		logger.Log.Error("create user error", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully", "data": res})
}
func (h *ClientHandler) Login(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Log.Error("binding error", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client, err := h.service.Login(context.Background(), &request)
	if err != nil {
		logger.Log.Error("login error", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User logged in successfully", "data": client})
}
func (h *ClientHandler) Refresh(c *gin.Context) {
	Bearer := "Bearer "
	authHeader := c.GetHeader("Authorization")
	strToken := authHeader[len(Bearer):]
	res, err := h.service.Refresh(context.Background(), strToken)
	if err != nil {
		logger.Log.Error("refresh error", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User refreshed successfully", "data": res})
}
func (h *ClientHandler) Register(c *gin.Context) {
	// get data
	var req RegisterReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		logger.Log.Error("felid to bind req data ", "error", err)
		c.JSON(200, gin.H{"message": "Not valid request body"})
	}
	// h.service.CreateAssessment()
	// check toke
	//TODO  check client data on db
	var registerRes RegisterRes
	registerRes.UserType, err = h.service.CheckUser(context.Background(), req.Email)
	if err != nil {
		slog.Error("failed to check Client type ", "err", err)
		shared.Respond(c, nil, &shared.AppError{Error: err})
		return
	}
	switch registerRes.UserType {
	case USER_TYPE_REGISTERED:
		slog.Debug("user already registered", "clientType", registerRes.UserType, "email", req.Email)
		shared.Respond(c, registerRes, &shared.AppError{})
		// return
	case USER_TYPE_EMAIL_NOT_VERIFIED:
		slog.Debug("email not verified", "clientType", registerRes, "email", req.Email)
		shared.Respond(c, registerRes.UserType, &shared.AppError{Error: err})
		// return
	case USER_TYPE_NOT_REGISTERED:

	}
	// if exist
	// 		if email verified
	// 			type: Registered
	// 		if not email not verified
	// 			type: EMAIL_NOT_VERIFIED
	// 			err=Email not verified
	// 			nav to email verify page
	// if not exist
	// 		type:NEW_USER
	// 		Nav to create user
}
