package sheard

import "github.com/golang-jwt/jwt/v5"

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
