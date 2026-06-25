package client

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	recaptchapb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	customErrors "github.com/AhmedZeyad/Akalni/customErrors"

	"github.com/AhmedZeyad/Akalni/middleware"
	"golang.org/x/crypto/bcrypt"
)

type ClientService struct {
	client     ClientRepo
	jwtService *middleware.JWTService
}

func NewClientService(client ClientRepo, jwtService *middleware.JWTService) *ClientService {
	return &ClientService{client: client,
		jwtService: jwtService,
	}
}

func (s ClientService) GetProfile(id int64) (client Client, err error) {
	client, err = s.client.GetByID(id)
	if err != nil {
		return client, err
	}
	return client, nil
}
func (cs *ClientService) CreateUser(ctx context.Context, clientrequest *RegisterRequest) (res RegisterResponse, err error) {
	if clientrequest.Email == "" {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if clientrequest.Password == "" {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	if clientrequest.ConfirmPassword == "" {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_CONFIRM_PASSWORD_REQUIRED)
	}
	if clientrequest.Password != clientrequest.ConfirmPassword {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_PASSWORD_MISMATCH)
	}
	if clientrequest.FirstName == "" {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_FIRST_NAME_REQUIRED)
	}
	if clientrequest.LastName == "" {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_LAST_NAME_REQUIRED)
	}
	if clientrequest.PhoneNumber == "" {
		return RegisterResponse{}, errors.New(customErrors.VALIDATION_PHONE_REQUIRED)
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(clientrequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, err
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
		return RegisterResponse{}, err
	}
	res.Client = *client.ToResponse()
	// Todo gen token
	res.Token, err = cs.jwtService.ClientGenToken(middleware.User{
		ID:              client.ID,
		Name:            client.FirstName + " " + client.LastName,
		Email:           client.Email,
		IsEmailVerified: client.IsEmailVerified,
	})
	if err != nil {
		return RegisterResponse{}, err
	}
	res.RefreshToken, err = cs.jwtService.ClientGenRefreshToken(middleware.User{
		ID:              client.ID,
		Email:           client.Email,
		IsEmailVerified: client.IsEmailVerified,
	})
	if err != nil {
		return RegisterResponse{}, err
	}

	return res, nil
}
func (cs *ClientService) Login(ctx context.Context, req *LoginRequest) (res RegisterResponse, err error) {
	if req.Email == "" {
		return res, errors.New(customErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if req.Password == "" {
		return res, errors.New(customErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	client, err := cs.client.GetByEmail(req.Email)
	if err != nil && err.Error() != customErrors.AUTH_EMAIL_NOT_VERIFIED {

		return res, err
	}
	slog.Debug("pass log", "req", req.Password)
	slog.Debug("pass log", "real", client.Password)

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(req.Password))
	if err != nil {
		return res, errors.New(customErrors.AUTH_PASSWORD_INCORRECT)
	}

	res.Client = *client.ToResponse()
	// Todo gen token
	res.Token, err = cs.jwtService.ClientGenToken(middleware.User{
		ID:              client.ID,
		Name:            client.FirstName + " " + client.LastName,
		Email:           client.Email,
		IsEmailVerified: client.IsEmailVerified,
	})
	if err != nil {
		return RegisterResponse{}, err
	}
	res.RefreshToken, err = cs.jwtService.ClientGenRefreshToken(middleware.User{
		ID:              client.ID,
		Email:           client.Email,
		IsEmailVerified: client.IsEmailVerified,
	})
	if err != nil {
		return RegisterResponse{}, err
	}

	return res, nil
}
func (cs *ClientService) Refresh(ctx context.Context, token string) (res RefreshTokenRes, err error) {

	cliams, err := cs.jwtService.UserTokenEvaluation(token, middleware.EvalRefreshToken)
	if err != nil {
		slog.Error("failed to verify refresh token", "error", err)
		return res, err
	}

	client, err := cs.client.GetByID(cliams.ID)
	if err != nil {
		slog.Error("failed to get by id", "id", cliams.ID, "error", err)
		return res, err
	}
	// Todo gen token
	res.Token, err = cs.jwtService.ClientGenToken(middleware.User{
		ID:              client.ID,
		Name:            client.FirstName + " " + client.LastName,
		Email:           client.Email,
		IsEmailVerified: client.IsEmailVerified,
	})
	if err != nil {
		slog.Error("failed to generate token", "error", err)
		return
	}

	return res, nil
}
func (cs *ClientService) CheckUser(ctx context.Context, email string) (string, error) {
	user, err := cs.client.GetByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return USER_TYPE_NOT_REGISTERED, nil
		}
		return "", err
	}
	return user.CheckUserType(), nil

}
func (cs *ClientService) UpdateEmail(clientID int64, email string) error {
	return cs.client.UpdateEmail(clientID, email)
}
func (cs *ClientService) SendEmailVerification(from, key, email string, otpLenght, otpExpire, salt int, otpType OTPType) error {

	client, err := cs.client.GetByEmail(email)
	if err != nil && (otpType == OTP_TYPE_EMAIL_VERIFICATION && err.Error() != customErrors.AUTH_EMAIL_NOT_VERIFIED) {
		slog.Error("failed to get client by email", "error", err)
		return err
	}

	otp, err := generateOtp(otpLenght)
	if err != nil {
		slog.Error("failed to generate otp", "error", err)
		return err
	}
	hashOtp := Hash(otp, salt)
	err = cs.client.SetOtpCode(OTPVerification{ClientID: client.ID, Code: hashOtp, Type: otpType}, otpExpire)
	if err != nil {
		slog.Error("failed to set otp code", "error", err)
		return err
	}
	err = middleware.SendOTP(from, key, client.Email, otp, otpType)
	if err != nil {
		slog.Error("failed to send otp", "error", err)
		return err
	}
	return nil
}
func (cs *ClientService) VerifyEmail(clientID int64, req OTPVerificationRequest, otpSalt int, otpType OTPType) error {

	client, err := cs.client.GetByID(clientID)
	if err != nil && err.Error() != customErrors.AUTH_EMAIL_NOT_VERIFIED {
		slog.Error("failed to get client by id", "error", err, "clientID", clientID)
		return err
	}

	otp, err := cs.client.GetOtpCode(clientID, otpType)
	if err != nil {
		slog.Error("failed to get otp code", "error", err)
		return err
	}
	if !checkOtp(otp.Code, req.Code, otpSalt) {
		err = fmt.Errorf("invalid otp")
		slog.Error("failed to verify otp", "error", err)
		return err
	}
	err = cs.client.VerifyOtpCode(otp.ID, clientID, otpType)
	if err != nil {
		slog.Error("failed to set otp code", "error", err)
		return err
	}
	switch otpType {
	case OTP_TYPE_EMAIL_VERIFICATION:
		err = cs.client.VerifyOtpCode(otp.ID, clientID, otpType)
		if err != nil {
			slog.Error("failed to verify otp", "error", err)
			return err
		}
		err = cs.client.UpdateClientEmailVerified(clientID)
		if err != nil {
			slog.Error("failed to update client email verified", "error", err)
			return err
		}
	case OTP_TYPE_EMAIL_UPDATE:
		err = cs.UpdateEmail(client.ID, req.Email)
		if err != nil {
			slog.Error("failed to update email", "error", err)
			return err
		}
	}
	return nil
}

// func (cs *ClientService)
func (cs *ClientService) CreateAssessment(projectID string, recaptchaKey string, token string, recaptchaAction string) {

	// Create the reCAPTCHA client.

	ctx := context.Background()
	client, err := recaptcha.NewClient(ctx)
	if err != nil {
		fmt.Printf("Error creating reCAPTCHA client\n")
	}
	defer client.Close()

	// Set the properties of the event to be tracked.
	event := &recaptchapb.Event{
		Token:   token,
		SiteKey: recaptchaKey,
	}

	assessment := &recaptchapb.Assessment{
		Event: event,
	}

	// Build the assessment request.
	request := &recaptchapb.CreateAssessmentRequest{
		Assessment: assessment,
		Parent:     fmt.Sprintf("projects/%s", projectID),
	}

	response, err := client.CreateAssessment(
		ctx,
		request)

	if err != nil {
		fmt.Printf("Error calling CreateAssessment: %v", err.Error())
	}

	// Check if the token is valid.
	if !response.TokenProperties.Valid {
		fmt.Printf("The CreateAssessment() call failed because the token was invalid for the following reasons: %v",
			response.TokenProperties.InvalidReason)
		return
	}

	// Check if the expected action was executed.
	if response.TokenProperties.Action != recaptchaAction {
		fmt.Printf("The action attribute in your reCAPTCHA tag does not match the action you are expecting to score")
		return
	}

	// Get the risk score and the reason(s).
	// For more information on interpreting the assessment, see:
	// https://cloud.google.com/recaptcha/docs/interpret-assessment
	fmt.Printf("The reCAPTCHA score for this token is:  %v", response.RiskAnalysis.Score)

	for _, reason := range response.RiskAnalysis.Reasons {
		fmt.Printf(reason.String() + "\n")
	}
}
