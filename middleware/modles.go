package middleware

import (
	"errors"
	"log/slog"
	"slices"
	"time"

	"github.com/AhmedZeyad/Akalni/customErrors"
	"github.com/AhmedZeyad/Akalni/utils"
	"github.com/golang-jwt/jwt/v5"
)

type AdminClaims struct {
	Email string `json:"email"`
	// Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func (admin *AdminClaims) ClaimsEval(evalType EvalClaimsType) error {
	switch evalType {
	case EvalToken:
		if utils.IsEmpty(admin.Subject) || admin.Subject != "DashBoard" {
			slog.Error("admin subject is not set or invalid", "subject", admin.Subject)
			return errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
		}

	case EvalRefreshToken:
		if utils.IsEmpty(admin.Subject) || admin.Subject != "Refresh" {
			slog.Error("admin subject is not set or invalid", "subject", admin.Subject)
			return errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
		}

	}
	if utils.IsEmpty(admin.ExpiresAt) || admin.ExpiresAt.Before(time.Now()) {
		slog.Error("admin token is expired")
		return errors.New(customErrors.AUTH_TOKEN_EXPIRED)
	}
	if utils.IsEmpty(admin.ID) {
		slog.Error("admin id is not set", "id", admin.ID)
		return errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
	}
	if utils.IsEmpty(admin.Email) {
		slog.Error("admin email is not set", "email", admin.Email)
		return errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
	}
	if utils.IsEmpty(admin.Issuer) || admin.Issuer != "Akalni" {
		slog.Error("admin issuer is not set or invalid", "issuer", admin.Issuer)
		return errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
	}

	if utils.IsEmpty(admin.Audience...) || !slices.Contains(admin.Audience, "admin") {
		slog.Error("admin audience is not set or invalid", "audience", admin.Audience)
		return errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
	}
	return nil
}

type JWTService struct {
	JWTExpire        int
	RefreshJWTExpire int
	JWTSecret        string
}

// admin
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Claims interface {
	jwt.Claims
	ClaimsEval(EvalClaimsType) error
}
type EvalClaimsType string

const (
	EvalToken        EvalClaimsType = "Token"
	EvalRefreshToken EvalClaimsType = "RefreshToken"
)
