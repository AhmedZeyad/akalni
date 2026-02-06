package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ba7rIbrahim/Akalni/customeErrors"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: Implement authentication middleware
// TODO: CLIMS STRUCT
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

func NewJWTService(jwtExpier, REfreshJwtExpier int, secret string) *JTWSevice {
	return &JTWSevice{
		JWTExpire:        jwtExpier,
		RefreshJWTExpire: REfreshJwtExpier,
		JWTSecret:        secret,
	}
}

// TODO: GEN JWT
func (jwtservice *JTWSevice) TokenGenrate(client Client) (stringToken string, err error) {
	claims := Claims{
		Email:           client.Email,
		IsEmailVerified: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprintf("%d", client.ID),
			Issuer:    "Akalni",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(jwtservice.JWTExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "client",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	stringToken, err = token.SignedString([]byte(jwtservice.JWTSecret))
	if err != nil {
		log.Printf("error on generate token: %v", err)
		return "", err
	}
	return stringToken, nil
}
func (jwtservice *JTWSevice) GenRefreshToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ID:        fmt.Sprintf("%d", id),
		Issuer:    "Akalni",
		Subject:   "Refresh",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(jwtservice.RefreshJWTExpire))),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})
	stringToken, err := token.SignedString([]byte(jwtservice.JWTSecret))
	if err != nil {
		log.Printf("error on generate token: %v", err)
		return "", err
	}
	return stringToken, nil
}

// TODO: VERIFY JWT
func (jwtservice *JTWSevice) TokenVerify(stringtToken string) error {
	token, err := jwt.ParseWithClaims(stringtToken, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(jwtservice.JWTSecret), nil
	})
	if err != nil || token == nil || !token.Valid {
		log.Printf("error on verify token: %v", err)
		return errors.New("AUTH_INVALID_TOKEN")
	}
	if token.Method != jwt.SigningMethodHS512 {
		return errors.New("AUTH_INVALID_TOKEN")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return errors.New(customeErrors.AUTH_INVALID_CREDENTIALS)
	}
	if claims.Email == "" {
		return errors.New(customeErrors.AUTH_INVALID_CREDENTIALS)
	}
	if claims.Subject != "client" {
		return errors.New(customeErrors.AUTH_INVALID_CREDENTIALS)
	}
	if claims.Issuer != "Akalni" {
		return errors.New(customeErrors.AUTH_INVALID_CREDENTIALS)
	}
	// TODO :implement lock account if emali not verified affter 1 month from creations
	return nil
}

// TODO: AUTH MIDDLEWARE
