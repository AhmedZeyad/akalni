package client

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/AhmedZeyad/Akalni/config"
	"github.com/AhmedZeyad/Akalni/logger"
	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service ClientService
	conf    config.Config
}

func NewClientHandler(service ClientService, conf config.Config) *ClientHandler {
	return &ClientHandler{service: service, conf: conf}
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
	case USER_TYPE_EMAIL_NOT_VERIFIED, USER_TYPE_NOT_REGISTERED:
		slog.Debug("email not verified", "clientType", registerRes, "email", req.Email)
		// generate otp
		otp, err := generateOtp(h.conf.OtpLenght)
		if err != nil {
			slog.Error("failed to generate otp", "error", err)
			shared.Respond(c, nil, &shared.AppError{Error: err})
			return
		}
		// save otp in db
		// hasPassword := sha256.Sum256([]byte(otp + conf.Config.OTPSalt))

		// send otp to email it not dev
		//

		// return  code on dev
		if h.conf.ISDev == "true" {
			shared.Respond(c, otp, &shared.AppError{})
			return
		}
		shared.Respond(c, registerRes.UserType, &shared.AppError{Error: err})
		// return
		// case USER_TYPE_NOT_REGISTERED:

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

func (ch *ClientHandler) SendOtp(c *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var req OTPVerificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	if req.Email == "" {
		shared.Respond(c, nil, &shared.AppError{Message: "email is required"})
		return
	}
	err = ch.service.SendEmailVerification(
		ch.conf.OtpEmailSender,
		ch.conf.OtpAppPassword,
		req.Email,
		ch.conf.OtpLenght,
		ch.conf.OtpExpire,
		ch.conf.OTPSalt,
		OTP_TYPE_EMAIL_VERIFICATION,
	)
	if err != nil {
		appError.Error = err
		slog.Error("failed to send email verification", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	shared.Respond(c, nil, nil)
}

func (ch *ClientHandler) ResendOtp(c *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var req OTPVerificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	err = ch.service.SendEmailVerification(
		ch.conf.OtpEmailSender,
		ch.conf.OtpAppPassword,
		req.Email,
		ch.conf.OtpLenght,
		ch.conf.OtpExpire,
		ch.conf.OTPSalt,
		OTP_TYPE_EMAIL_VERIFICATION,
	)
	if err != nil {
		appError.Error = err
		slog.Error("failed to send email verification", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	shared.Respond(c, nil, nil)
}

func (ch *ClientHandler) VerifyOtp(c *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var req OTPVerificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	err = ch.service.VerifyEmail(req, ch.conf.OTPSalt, OTP_TYPE_EMAIL_VERIFICATION)
	if err != nil {
		appError.Error = err
		slog.Error("failed to verify email", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	shared.Respond(c, nil, nil)
}

// auth
//
//	func (ch *ClientHandler) SendEmailOTP(c *gin.Context) {
//		appError := &shared.AppError{Layer: "Handler"}
//		var req UpdateEmailRequest
//		if err := c.ShouldBindJSON(&req); err != nil {
//			appError.Error = err
//			slog.Error("failed to bind json", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		otp, err := generateOtp(ch.conf.OtpLenght)
//		if err != nil {
//			appError.Error = err
//			slog.Error("failed to generate otp", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		client := c.MustGet("client").(Client)
//		appError.Error = ch.service.client.SetOtpCode(OTPVerification{
//			Code: otp, ClientID: client.ID,
//			Type: OTP_TYPE_EMAIL_VERIFICATION,
//		}, ch.conf.OtpExpire)
//		if appError.Error != nil {
//			slog.Error("failed to update email", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		shared.Respond(c, nil, nil)
//	}
//
//	func (ch *ClientHandler) SendOtpForUpdateEmail(c *gin.Context) {
//		appError := &shared.AppError{Layer: "Handler"}
//		var req UpdateEmailRequest
//		if err := c.ShouldBindJSON(&req); err != nil {
//			appError.Error = err
//			slog.Error("failed to bind json", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		client, err := ch.service.client.GetByEmail(req.Email)
//		if err != nil {
//			appError.Error = err
//			slog.Error("failed to get client by email", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		if client.Email == req.Email {
//			appError.Error = fmt.Errorf("email cannot be updated to the same email")
//			shared.Respond(c, nil, appError)
//			return
//		}
//		otp, err := generateOtp(ch.conf.OtpLenght)
//		if err != nil {
//			appError.Error = err
//			slog.Error("failed to generate otp", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		appError.Error = ch.service.client.SetOtpCode(OTPVerification{
//			Code: otp, ClientID: client.ID,
//			Type: OTP_TYPE_EMAIL_UPDATE,
//		}, ch.conf.OtpExpire)
//		if appError.Error != nil {
//			slog.Error("failed to update email", "error", appError)
//			shared.Respond(c, nil, appError)
//			return
//		}
//		shared.Respond(c, otp, nil)
//	}
// func (ch *ClientHandler) VerifyUpdateEmail(c *gin.Context) {
// 	appError := &shared.AppError{Layer: "Handler"}
// 	var req UpdateEmailRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		appError.Error = err
// 		slog.Error("failed to bind json", "error", appError)
// 		shared.Respond(c, nil, appError)
// 		return
// 	}
// 	otp, err := generateOtp(ch.conf.OtpLenght)
// 	if err != nil {
// 		appError.Error = err
// 		slog.Error("failed to generate otp", "error", appError)
// 		shared.Respond(c, nil, appError)
// 		return
// 	}
// 	if !checkOtp(otp, req.Code, ch.conf.OTPSalt) {
// 		appError.Error = fmt.Errorf("invalid otp")
// 		shared.Respond(c, nil, appError)
// 		return
// 	}

// 	shared.Respond(c, otp, nil)
// }

func (ch *ClientHandler) SendOtpForUpdateEmail(c *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var req OTPVerificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	email := req.Email
	if email == "" {
		shared.Respond(c, nil, &shared.AppError{Message: "email is required"})
		return
	}
	err = ch.service.SendEmailVerification(
		ch.conf.OtpEmailSender,
		ch.conf.OtpAppPassword,
		email,
		ch.conf.OtpLenght,
		ch.conf.OtpExpire,
		ch.conf.OTPSalt,
		OTP_TYPE_EMAIL_UPDATE,
	)
	if err != nil {
		appError.Error = err
		slog.Error("failed to send email verification", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	shared.Respond(c, nil, nil)
}
func (ch *ClientHandler) ResendOtpForUpdateEmail(c *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var req OTPVerificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	email := req.Email
	if email == "" {
		shared.Respond(c, nil, &shared.AppError{Message: "email is required"})
		return
	}
	err = ch.service.SendEmailVerification(
		ch.conf.OtpEmailSender,
		ch.conf.OtpAppPassword,
		email,
		ch.conf.OtpLenght,
		ch.conf.OtpExpire,
		ch.conf.OTPSalt,
		OTP_TYPE_EMAIL_UPDATE,
	)
	if err != nil {
		appError.Error = err
		slog.Error("failed to send email verification", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	shared.Respond(c, nil, nil)
}

func (ch *ClientHandler) VerifyUpdateEmail(c *gin.Context) {
	appError := &shared.AppError{Layer: "Handler"}
	var req OTPVerificationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		appError.Error = err
		slog.Error("failed to bind json", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	err = ch.service.VerifyEmail(req, ch.conf.OTPSalt, OTP_TYPE_EMAIL_UPDATE)
	if err != nil {
		appError.Error = err
		slog.Error("failed to verify email", "error", appError)
		shared.Respond(c, nil, appError)
		return
	}
	shared.Respond(c, nil, nil)
}

func generateOtp(length int) (string, error) {
	// const digits = "0123456789"
	// otp := make([]byte, length)
	number, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", number.Int64()+100000), nil

}
func Hash(otp string, salt int) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", otp, salt)))
	return fmt.Sprintf("%x", hash)
}
func checkOtp(otp, reqOTP string, salt int) bool {
	reqOTP = Hash(reqOTP, salt)

	return otp == reqOTP
}
