package users

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service UserService
	jwt     *middleware.JWTService
}

func NewUserHandler(service UserService, jwt *middleware.JWTService) *UserHandler {
	return &UserHandler{service: service, jwt: jwt}
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var userFromRequest UserFromRequest
	appError.Error = ctx.ShouldBindJSON(&userFromRequest)
	if appError.Error != nil {
		slog.Error("failed to bind json", "error", appError.Error)
		shared.Respond(ctx, nil, appError)
		return
	}
	appError.Error = uh.service.CreateUser(userFromRequest)
	if appError.Error != nil {
		shared.Respond(ctx, nil, appError)
		return
	}
	shared.Respond(ctx, "nil", nil)
}
func (uh *UserHandler) Login(ctx *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var userFromRequest UserFromRequest
	appError.Error = ctx.ShouldBindJSON(&userFromRequest)
	if appError.Error != nil {
		slog.Error("failed to bind json", "error", appError.Error)
		shared.Respond(ctx, nil, appError)
		return
	}
	userFromRequest.Validate("login")
	user, err := uh.service.Login(userFromRequest)
	if err != nil {
		appError.Error = err
		shared.Respond(ctx, nil, appError)
		return
	}
	tokenUser := middleware.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	token, err := uh.jwt.UserGenToken(tokenUser)
	if err != nil {
		appError.Error = err
		shared.Respond(ctx, nil, appError)
		return
	}
	refreshToken, err := uh.jwt.UserGenRefreshToken(tokenUser)
	if err != nil {
		appError.Error = err
		shared.Respond(ctx, nil, appError)
		return
	}

	shared.Respond(ctx, user.ToUserResponse(token, refreshToken), nil)
}
func (uh *UserHandler) ResetPassword(ctx *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var resetPasswordRequest ResetPasswordRequest
	appError.Error = ctx.ShouldBindJSON(&resetPasswordRequest)
	if appError.Error != nil {
		slog.Error("failed to bind json", "error", appError.Error)
		shared.Respond(ctx, nil, appError)
		return
	}
	appError.Error = resetPasswordRequest.Validate()
	if appError.Error != nil {
		slog.Error("failed to validate", "error", appError.Error)
		shared.Respond(ctx, nil, appError)
		return
	}
	appError.Error = uh.service.ResetPassword(resetPasswordRequest)
	if appError.Error != nil {
		slog.Error("failed to reset password", "error", appError.Error)
		shared.Respond(ctx, nil, appError)
		return
	}
	shared.Respond(ctx, nil, nil)
}
