package client

import "time"

type Client struct {
	ID              int64      `db:"id" json:"id"`
	FirstName       string     `db:"first_name" json:"first_name"`
	LastName        string     `db:"last_name" json:"last_name"`
	Email           string     `db:"email" json:"email"`
	PhoneNumber     string     `db:"phone_number" json:"phone_number"`
	Password        string     `db:"password" json:"password"`
	ConfirmPassword string     `db:"confirm_password" json:"confirm_password"`
	IsEmailVerified bool       `db:"is_email_verified" json:"is_email_verified"`
	EmailVerifiedAt *time.Time `db:"email_verified_at" json:"email_verified_at"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"-"`
	UpdatedBy       int        `db:"updated_by" json:"-"`
	DeletedAt       time.Time  `db:"deleted_at" json:"-"`
	DeletedBy       int        `db:"deleted_by" json:"-"`
}

type ClientResponse struct {
	ID              int64      `json:"id"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Email           string     `json:"email"`
	PhoneNumber     string     `json:"phone_number"`
	IsEmailVerified bool       `json:"is_email_verified"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
}
type RegisterResponse struct {
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	Client       ClientResponse `json:"client"`
}
type RegisterRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RefreshTokenRes struct {
	Token string `json:"token"`
}
type RegisterReq struct {
	Stoken string `json:"stoke"`
	Email  string `json:"email"`
}
type RegisterRes struct {
	UserType string `json:"user_type"`
}

const (
	USER_TYPE_REGISTERED         = "REGISTERED"
	USER_TYPE_EMAIL_NOT_VERIFIED = "EMAIL_NOT_VERIFIED"
	USER_TYPE_NOT_REGISTERED     = "NOT_REGISTERED"
	USER_TYPE_BLOCKED            = "BLOCKED"
)

func (client *Client) CheckUserType() string {
	if client.EmailVerifiedAt == nil && !client.IsEmailVerified {
		return USER_TYPE_EMAIL_NOT_VERIFIED
	}
	// TODO implement user blocking account

	// if client.IsBlocked {
	// 	return USER_TYPE_BLOCKED
	// }
	return USER_TYPE_REGISTERED
}

type OTPVerification struct {
	ID        int64     `db:"id" json:"id"`
	ClientID  int64     `db:"client_id" json:"client_id"`
	Code      string    `db:"otp_code" json:"otp_code"`
	Type      OTPType   `db:"type" json:"type"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
}

type OTPVerificationRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
type UpdateEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
type OTPType string

const (
	OTP_TYPE_EMAIL_UPDATE       OTPType = "EMAIL_UPDATE"
	OTP_TYPE_EMAIL_VERIFICATION OTPType = "EMAIL_VERIFICATION"
)

func (c *Client) ToResponse() *ClientResponse {
	return &ClientResponse{
		ID:              c.ID,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		Email:           c.Email,
		PhoneNumber:     c.PhoneNumber,
		IsEmailVerified: c.IsEmailVerified,
		EmailVerifiedAt: c.EmailVerifiedAt,
	}
}
