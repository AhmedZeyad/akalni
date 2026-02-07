package auth

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"github.com/ba7rIbrahim/Akalni/customeErrors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	client     AuthRepo
	jwtService *JTWSevice
}

func NewAuthService(client AuthRepo, jwtService *JTWSevice) *AuthService {
	return &AuthService{
		client:     client,
		jwtService: jwtService,
	}
}

func (cs *AuthService) CreateUser(ctx context.Context, clientrequest *RegisterRequest) (res RegisterRespons, err error) {
	if clientrequest.Email == "" {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if clientrequest.Password == "" {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	if clientrequest.ConfirmPassword == "" {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_CONFIRM_PASSWORD_REQUIRED)
	}
	if clientrequest.Password != clientrequest.ConfirmPassword {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_PASSWORD_MISMATCH)
	}
	if clientrequest.FirstName == "" {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_FIRST_NAME_REQUIRED)
	}
	if clientrequest.LastName == "" {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_LAST_NAME_REQUIRED)
	}
	if clientrequest.PhoneNumber == "" {
		return RegisterRespons{}, errors.New(customeErrors.VALIDATION_PHONE_REQUIRED)
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(clientrequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterRespons{}, err
	}
	client := &Client{
		FirstName:   clientrequest.FirstName,
		LastName:    clientrequest.LastName,
		PhoneNumber: clientrequest.PhoneNumber,
		Email:       clientrequest.Email,
		Password:    string(pass),
	}

	err = cs.client.Create(ctx, client)
	if err != nil {
		return RegisterRespons{}, err
	}
	res.Client = ClientReqponse{
		ID:          client.ID,
		FirstName:   client.FirstName,
		LastName:    client.LastName,
		PhoneNumber: client.PhoneNumber,
		Email:       client.Email,
	}
	// Todo gen token
	res.Token, err = cs.jwtService.TokenGenrate(*client)
	if err != nil {
		return RegisterRespons{}, err
	}
	res.RefreshToken, err = cs.jwtService.GenRefreshToken(client.ID)
	if err != nil {
		return RegisterRespons{}, err
	}

	return res, nil
}
func (cs *AuthService) Login(ctx context.Context, req *LoginRequest) (res RegisterRespons, err error) {
	if req.Email == "" {
		return res, errors.New(customeErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if req.Password == "" {
		return res, errors.New(customeErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	client, err := cs.client.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}
	slog.Error("pass log", "req", req.Password)
	slog.Error("pass log", "real", client.Password)

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(req.Password))
	if err != nil {
		return res, errors.New(customeErrors.AUTH_PASSWORD_INCORRECT)
	}
	slog.Error("pass log", "req", req.Password)

	res.Client =
		// res=
		ClientReqponse{
			ID:          client.ID,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			PhoneNumber: client.PhoneNumber,
			Email:       client.Email,
		}
	// Todo gen token
	res.Token, err = cs.jwtService.TokenGenrate(client)
	if err != nil {
		return RegisterRespons{}, err
	}
	res.RefreshToken, err = cs.jwtService.GenRefreshToken(client.ID)
	if err != nil {
		return RegisterRespons{}, err
	}

	return res, nil
}
func (cs *AuthService) Refresh(ctx context.Context, token string) (res RefreshTokenRes, err error) {

	cliams, err := cs.jwtService.RefreshTokenVerify(token)
	if err != nil {
		slog.Error("failed to verify refresh token", "error", err)
		return res, err
	}
	id, err := strconv.Atoi(cliams.ID)
	if err != nil {
		slog.Error("failed to convert  to int", "id", cliams.ID, "error", err)
		return res, err
	}
	client, err := cs.client.GetByID(ctx, id)
	if err != nil {
		slog.Error("failed to get by id", "id", id, "error", err)
		return res, err
	}
	// Todo gen token
	res.Token, err = cs.jwtService.TokenGenrate(client)
	if err != nil {
		slog.Error("failed to generate token", "error", err)
		return
	}

	return res, nil
}
