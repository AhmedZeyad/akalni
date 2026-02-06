package auth

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ba7rIbrahim/Akalni/customeErrors"
	"golang.org/x/crypto/bcrypt"
)

type ClientService struct {
	client     AuthRepo
	jwtService *JTWSevice
}

func NewClientService(client AuthRepo, jwtService *JTWSevice) *ClientService {
	return &ClientService{
		client:     client,
		jwtService: jwtService,
	}
}

func (cs *ClientService) CreateUser(ctx context.Context, clientrequest *RegisterRequest) (res RegisterRespons, err error) {
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
func (cs *ClientService) Login(ctx context.Context, req *LoginRequest) (res RegisterRespons, err error) {
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
