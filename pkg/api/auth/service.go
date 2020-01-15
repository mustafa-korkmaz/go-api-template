package auth

import (
	"strings"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/uow"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/helper"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/role"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/secure"

	"github.com/labstack/echo"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/request"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/response"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service represents auth api interface
type Service interface {
	Authenticate(string, string) *response.APIResponse
	Register(*request.Register) *response.APIResponse
	Refresh(echo.Context) *response.APIResponse
}

//Uow represents olive service unit of work behaviour
type Uow interface {
	Save() error
	//begin transaction
}

// Auth represents auth api service
type Auth struct {
	uow            Uow
	repository     Repository
	tokenGenerator TokenGenerator
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateTokens(*model.User) (string, string, error)
}

// Repository represents user auth operations interface
type Repository interface {
	GetUserByEmail(email string) *model.User
	GetUserByUsername(username string) *model.User
	Register(*model.User) error
	FindOneByID(id string) *mongo.SingleResult
}

// Authenticate tries to authenticate user
func (a *Auth) Authenticate(usernameOrEmail string, password string) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
		Data:   nil,
	}

	var checkByEmail = false
	if strings.ContainsAny(usernameOrEmail, "@") {
		checkByEmail = true
	}

	var user = new(model.User)

	if checkByEmail {
		user = a.repository.GetUserByEmail(usernameOrEmail)
	} else {
		user = a.repository.GetUserByUsername(usernameOrEmail)
	}

	if user == nil {
		apiResp.SetError(enum.ErrorCode.UserNotFound)
		return &apiResp
	}

	isPasswordValid := secure.ValidatePassword(user.Password, password)

	if !isPasswordValid {
		apiResp.SetError(enum.ErrorCode.UserNotAuthorized)
		return &apiResp
	}

	//password is valid now create access tokens
	var token, refreshToken, err = a.tokenGenerator.GenerateTokens(user)

	if err != nil {
		apiResp.SetError(enum.ErrorCode.TokenCreationError)
		return &apiResp
	}

	apiResp.Data = response.Auth{
		Email:        user.Email,
		ID:           user.ID.Hex(),
		RefreshToken: refreshToken,
		Token:        token,
	}

	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// Refresh returns a new valid token and a refresh token
func (a *Auth) Refresh(c echo.Context) *response.APIResponse {

	var (
		err          error
		token        string
		refreshToken string
	)

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	var user model.User

	var id = role.User(c).ID

	var doc = a.repository.FindOneByID(id)

	err = doc.Decode(&user)

	if err != nil {
		apiResp.SetError(enum.ErrorCode.UserNotFound)
		return &apiResp
	}

	token, refreshToken, err = a.tokenGenerator.GenerateTokens(&user)

	if err != nil {
		apiResp.SetError(enum.ErrorCode.TokenCreationError)
		return &apiResp
	}

	apiResp.Data = response.Auth{
		Email:        user.Email,
		ID:           user.ID.Hex(),
		RefreshToken: refreshToken,
		Token:        token,
	}

	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// Register creates a new user
func (a *Auth) Register(registerRequest *request.Register) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	errorCode := a.validateNewRegistration(registerRequest)

	if errorCode != enum.ErrorCode.None {
		apiResp.SetError(errorCode)
		return &apiResp
	}

	passwordHash := secure.Hash(registerRequest.Password)

	//create user
	user := &model.User{
		Username:    registerRequest.Username,
		Email:       registerRequest.Email,
		FullName:    registerRequest.FullName,
		AccessLevel: enum.AccessLevel.Standart,
		Status:      enum.UserStatus.Created,
		Password:    passwordHash,
		CreatedAt:   helper.UtcNow(),
	}

	var err = a.repository.Register(user)

	if err != nil {
		apiResp.SetError(enum.ErrorCode.AppError)
		return &apiResp
	}

	apiResp.Data = user.ID.Hex()
	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// New creates new auth api service
func New(tg TokenGenerator, client *mongo.Client, dbName string) *Auth {

	var uow = uow.New(client, dbName)

	var auth = Auth{}
	auth.repository = uow.GetUserRepository()
	auth.uow = uow
	auth.tokenGenerator = tg
	return &auth
}

func (a *Auth) validateNewRegistration(registerRequest *request.Register) enum.ErrorCodeType {
	//check password strength
	var passStrong = secure.IsPasswordSecure(registerRequest.Password)

	if !passStrong {
		return enum.ErrorCode.WeakPassword
	}

	//check existance by email
	var user = a.repository.GetUserByEmail(registerRequest.Email)

	if user != nil {
		return enum.ErrorCode.UserAlreadyExists
	}

	//check existance by username
	user = a.repository.GetUserByUsername(registerRequest.Username)

	if user != nil {
		return enum.ErrorCode.UserAlreadyExists
	}

	return enum.ErrorCode.None
}
