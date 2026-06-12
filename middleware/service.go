package middleware

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJwtService(exp, refExp int, secret string) *JWTService {
	return &JWTService{
		JWTExpire:        exp,
		RefreshJWTExpire: refExp,
		JWTSecret:        secret,
	}
}

func (j *JWTService) genToken(claims jwt.Claims) (stringToken string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err = token.SignedString([]byte(j.JWTSecret))
	if err != nil {
		slog.Error("failed to get string token ", "error", err)
		return "", err
	}
	return
}
func (j *JWTService) UserGenToken(admin User) (stringToken string, err error) {
	claims := AdminClaims{
		ID:    admin.ID,
		Email: admin.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprint(admin.ID),
			Issuer:    "Akalni",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.JWTExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "DashBoard",
			Audience:  jwt.ClaimStrings{"admin"},
		},
	}
	token, err := j.genToken(claims)
	if err != nil {
		slog.Error("failed to generate admin token", "error", err)
		return "", err
	}
	return token, nil
}
func (j *JWTService) UserGenRefreshToken(user User) (string, error) {
	claims := AdminClaims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprint(user.ID),
			Issuer:    "Akalni",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.RefreshJWTExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Refresh",
			Audience:  jwt.ClaimStrings{"admin"},
		},
	}
	token, err := j.genToken(claims)
	if err != nil {
		slog.Error("failed to generate refresh token", "error", err)
		return "", err
	}
	return token, nil
}
func (j *JWTService) UserTokenEvaluation(strToken string, evalType EvalClaimsType) (AdminClaims, error) {
	claims, err := tokenEvaluation(strToken, j.JWTSecret, evalType, &AdminClaims{})
	if err != nil {
		slog.Error("failed to evaluate admin token", "error", err)
		return *claims, err
	}
	// TODO check user on db
	return *claims, nil
}

func (j *JWTService) ClientGenToken(client User) (stringToken string, err error) {
	claims := ClientClaims{
		ClientID:        client.ID,
		Name:            client.Name,
		IsEmailVerified: client.IsEmailVerified,
		Email:           client.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Akalni",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.JWTExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "ClientApp",
			Audience:  jwt.ClaimStrings{"client"},
		},
	}
	token, err := j.genToken(claims)
	if err != nil {
		slog.Error("failed to generate admin token", "error", err)
		return "", err
	}
	return token, nil
}
func (j *JWTService) ClientGenRefreshToken(client User) (string, error) {
	claims := ClientClaims{
		ClientID:        client.ID,
		Name:            client.Name,
		IsEmailVerified: client.IsEmailVerified,
		Email:           client.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Akalni",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.RefreshJWTExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Refresh",
			Audience:  jwt.ClaimStrings{"client"},
		},
	}
	token, err := j.genToken(claims)
	if err != nil {
		slog.Error("failed to generate refresh token", "error", err)
		return "", err
	}
	return token, nil
}
func (j *JWTService) ClientTokenEvaluation(strToken string, evalType EvalClaimsType) (ClientClaims, error) {
	claims, err := tokenEvaluation[*ClientClaims](strToken, j.JWTSecret, evalType, &ClientClaims{})
	if err != nil {
		slog.Error("failed to evaluate admin token", "error", err)
		return *claims, err
	}
	// TODO check user on db
	return *claims, nil
}

func tokenEvaluation[T Claims](strToken, secret string, evalType EvalClaimsType, claims T) (T, error) {
	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("unexpected signing method", "method", token.Method)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return []byte(secret), nil
	})
	if err != nil {
		slog.Error("failed to parse token", "error", err, "token", strToken)
		return claims, err
	}
	c, ok := token.Claims.(T)
	if !ok {
		slog.Error("invalid claims type")
		return claims, fmt.Errorf("invalid claims type")
	}
	if err := c.ClaimsEval(evalType); err != nil {
		slog.Error("claims evaluation failed", "error", err)
		return claims, err
	}
	return c, nil

}
