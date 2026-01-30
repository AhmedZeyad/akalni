package services

import (
	"context"
	"errors"

	"github.com/ba7rIbrahim/Akalni/Models"
	"github.com/ba7rIbrahim/Akalni/customeErrors"
	"github.com/ba7rIbrahim/Akalni/middlewares"
	"github.com/ba7rIbrahim/Akalni/repo"
	"golang.org/x/crypto/bcrypt"
)

type ClientService struct {
	client     repo.AuthRepo
	jwtService *middlewares.JTWSevice
}

func NewClientService(client repo.AuthRepo, jwtService *middlewares.JTWSevice) *ClientService {
	return &ClientService{
		client:     client,
		jwtService: jwtService,
	}
}

func (cs *ClientService) CreateUser(ctx context.Context, clientrequest *Models.RegisterRequest) (res Models.RegisterRespons, err error) {
	if clientrequest.Email == "" {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if clientrequest.Password == "" {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	if clientrequest.ConfirmPassword == "" {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_CONFIRM_PASSWORD_REQUIRED)
	}
	if clientrequest.Password != clientrequest.ConfirmPassword {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_PASSWORD_MISMATCH)
	}
	if clientrequest.FirstName == "" {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_FIRST_NAME_REQUIRED)
	}
	if clientrequest.LastName == "" {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_LAST_NAME_REQUIRED)
	}
	if clientrequest.PhoneNumber == "" {
		return Models.RegisterRespons{}, errors.New(customeErrors.VALIDATION_PHONE_REQUIRED)
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(clientrequest.Password), 4)
	if err != nil {
		return Models.RegisterRespons{}, err
	}
	client := &Models.Client{
		FirstName:   clientrequest.FirstName,
		LastName:    clientrequest.LastName,
		PhoneNumber: clientrequest.PhoneNumber,
		Email:       clientrequest.Email,
		Password:    string(pass),
	}

	err = cs.client.Create(ctx, client)
	if err != nil {
		return Models.RegisterRespons{}, err
	}
	res.Client = Models.ClientReqponse{
		ID:          client.ID,
		FirstName:   client.FirstName,
		LastName:    client.LastName,
		PhoneNumber: client.PhoneNumber,
		Email:       client.Email,
	}
	// Todo gen token
	res.Token, err = cs.jwtService.TokenGenrate(*client)
	if err != nil {
		return Models.RegisterRespons{}, err
	}
	res.RefreshToken, err = cs.jwtService.GenRefreshToken(client.ID)
	if err != nil {
		return Models.RegisterRespons{}, err
	}

	return res, nil
}
