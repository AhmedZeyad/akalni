package auth

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/AhmedZeyad/Akalni/customErrors"
// 	"github.com/golang-jwt/jwt/v5"
// )

// // TODO: Implement authentication middleware
// // TODO: CLIMS STRUCT
// type Claims struct {
// 	ClientID        int    `json:"client_id"`
// 	Email           string `json:"email"`
// 	IsEmailVerified bool   `json:"is_email_verified"`
// 	jwt.RegisteredClaims
// }
// type JTWSevice struct {
// 	JWTExpire        int
// 	RefreshJWTExpire int
// 	JWTSecret        string
// }

// func NewJWTService(jwtExpier, REfreshJwtExpier int, secret string) *JTWSevice {
// 	return &JTWSevice{
// 		JWTExpire:        jwtExpier,
// 		RefreshJWTExpire: REfreshJwtExpier,
// 		JWTSecret:        secret,
// 	}
// }

// // TODO: GEN JWT
// func (jwtservice *JTWSevice) TokenGenrate(client Client) (stringToken string, err error) {
// 	claims := Claims{
// 		ClientID:        client.ID,
// 		Email:           client.Email,
// 		IsEmailVerified: false,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    "Akalni",
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(jwtservice.JWTExpire))),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			Subject:   "client",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
// 	stringToken, err = token.SignedString([]byte(jwtservice.JWTSecret))
// 	if err != nil {
// 		log.Printf("error on generate token: %v", err)
// 		return "", err
// 	}
// 	return stringToken, nil
// }
// func (jwtservice *JTWSevice) GenRefreshToken(id int) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
// 		ID:        fmt.Sprintf("%d", id),
// 		Issuer:    "Akalni",
// 		Subject:   "Refresh",
// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(jwtservice.RefreshJWTExpire))),
// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
// 	})
// 	stringToken, err := token.SignedString([]byte(jwtservice.JWTSecret))
// 	if err != nil {
// 		log.Printf("error on generate token: %v", err)
// 		return "", err
// 	}
// 	return stringToken, nil
// }

// // TODO: VERIFY JWT
// func (jwtservice *JTWSevice) TokenVerify(stringtToken string) (Claims, error) {
// 	token, err := jwt.ParseWithClaims(stringtToken, &Claims{}, func(t *jwt.Token) (any, error) {
// 		return []byte(jwtservice.JWTSecret), nil
// 	})
// 	if err != nil || token == nil || !token.Valid {
// 		log.Printf("error on verify token: %v", err)
// 		return Claims{}, errors.New("AUTH_INVALID_TOKEN")
// 	}
// 	if token.Method != jwt.SigningMethodHS512 {
// 		return Claims{}, errors.New("AUTH_INVALID_TOKEN")
// 	}
// 	claims, ok := token.Claims.(*Claims)
// 	if !ok {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	if claims.ClientID == 0 {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	if claims.Email == "" {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	if claims.Subject != "client" {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	if claims.Issuer != "Akalni" {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	// TODO :implement lock account if emali not verified affter 1 month from creations
// 	return *claims, nil
// }

// func (jwtservice *JTWSevice) RefreshTokenVerify(strToken string) (claims Claims, err error) {
// 	token, err := jwt.ParseWithClaims(strToken, &claims, func(t *jwt.Token) (any, error) {
// 		return []byte(jwtservice.JWTSecret), nil
// 	})
// 	if err != nil {
// 		if errors.Is(err, jwt.ErrTokenExpired) {
// 			return Claims{}, errors.New(customErrors.AUTH_TOKEN_EXPIRED)
// 		}
// 		return Claims{}, err
// 	}
// 	if token == nil {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_TOKEN)
// 	}
// 	if !token.Valid {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_TOKEN)
// 	}
// 	if token.Method != jwt.SigningMethodHS512 {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_TOKEN)
// 	}
// 	if claims.ID == "" {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}

// 	if claims.Subject != "Refresh" {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	if claims.Issuer != "Akalni" {
// 		return Claims{}, errors.New(customErrors.AUTH_INVALID_CREDENTIALS)
// 	}
// 	// TODO :implement lock account if emali not verified affter 1 month from creations
// 	return claims, nil
// }
