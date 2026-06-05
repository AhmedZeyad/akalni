package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	recaptchapb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	customErrors "github.com/AhmedZeyad/Akalni/customErrors"
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
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if clientrequest.Password == "" {
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	if clientrequest.ConfirmPassword == "" {
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_CONFIRM_PASSWORD_REQUIRED)
	}
	if clientrequest.Password != clientrequest.ConfirmPassword {
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_PASSWORD_MISMATCH)
	}
	if clientrequest.FirstName == "" {
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_FIRST_NAME_REQUIRED)
	}
	if clientrequest.LastName == "" {
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_LAST_NAME_REQUIRED)
	}
	if clientrequest.PhoneNumber == "" {
		return RegisterRespons{}, errors.New(customErrors.VALIDATION_PHONE_REQUIRED)
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
		return res, errors.New(customErrors.VALIDATION_EMAIL_REQUIRED)
	}
	if req.Password == "" {
		return res, errors.New(customErrors.VALIDATION_PASSWORD_REQUIRED)
	}
	client, err := cs.client.GetByEmail(ctx, req.Email)
	if err != nil {

		return res, err
	}
	slog.Debug("pass log", "req", req.Password)
	slog.Debug("pass log", "real", client.Password)

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(req.Password))
	if err != nil {
		return res, errors.New(customErrors.AUTH_PASSWORD_INCORRECT)
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
func (cs *AuthService) CheckUser(ctx context.Context, email string) (string, error) {
	user, err := cs.client.GetByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return USER_TYPE_NOT_REGISTERED, nil
		}
		return "", err
	}
	return user.CheckUserType(), nil
}
func (cs *AuthService) CreateAssessment(projectID string, recaptchaKey string, token string, recaptchaAction string) {

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
