package utils

import (
	"log/slog"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func PassHash(password string) (string, error) {
	password = strings.ToLower(password[len(password)/2:]) + strings.ToUpper(password[:len(password)/2])
	bPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		return "nil", err
	}
	return string(bPass), nil
}
func ComparePass(password string, hashedPassword string) error {
	password = strings.ToLower(password[len(password)/2:]) + strings.ToUpper(password[:len(password)/2])

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		slog.Error("failed to compare password", "error", err)
		return err
	}
	return nil
}
func IsEmpty[T comparable](t ...T) bool {
	if len(t) == 0 {
		return true
	}
	if len(t) == 1 {
		var zero T
		if t[0] == zero {
			return true
		}
	}
	return false
}
