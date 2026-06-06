package users

import (
	"errors"
	"time"

	"github.com/AhmedZeyad/Akalni/customErrors"
	"github.com/AhmedZeyad/Akalni/utils"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy int64     `db:"created_by"`
}

func (u *User) ToUserResponse(token, refreshToken string) UserResponse {
	return UserResponse{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Status:       u.Status,
		Token:        token,
		RefreshToken: refreshToken,
	}
}

type UserFromRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserFromRequest) Validate(checkType string) error {
	if checkType != "login" {
		if utils.IsEmpty(u.Email) {
			return errors.New(customErrors.VALIDATION_EMAIL_REQUIRED)
		}
	}
	if utils.IsEmpty(u.Password) {
		return errors.New(customErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	// TODO add password validation like minimum length, special characters, etc.
	if len(u.Password) < 8 {
		return errors.New(customErrors.VALIDATION_PASSWORD_TOO_SHORT)
	}

	return nil

}

type UserResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       bool   `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (r *ResetPasswordRequest) Validate() error {
	if utils.IsEmpty(r.Email) {
		return errors.New(customErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if utils.IsEmpty(r.NewPassword) {
		return errors.New(customErrors.VALIDATION_CONFIRM_PASSWORD_REQUIRED)
	}
	if len(r.NewPassword) < 8 {
		return errors.New(customErrors.VALIDATION_PASSWORD_TOO_SHORT)
	}
	return nil
}
