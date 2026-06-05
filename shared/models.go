package shared

import (
	"github.com/AhmedZeyad/Akalni/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
	jwt.RegisteredClaims
}
type JTWSevice struct {
	JWTExpire        int
	RefreshJWTExpire int
	JWTSecret        string
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   *error `json:"error,omitempty"`
}
type AppError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Error   error  `json:"error"`
	Layer   string `json:"-"`
}

var Conf *config.Config
