package olive

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	u "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository/user"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/uow"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/response"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service represents auth api interface
type Service interface {
	Authenticate(string, string) *response.APIResponse
	Register() *response.APIResponse
	Refresh() *response.APIResponse
}

//Uow represents olive service unit of work behaviour
type Uow interface {
	Save() error
	//begin transaction
}

// Auth represents auth api service
type Auth struct {
	uow        Uow
	repository u.Repository
}

// Authenticate tries to authenticate user
func (a *Auth) Authenticate(usernameOrEmail string, password string) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
		Data:   nil,
	}

	// var olive model.Olive
	// var doc = a.repository.FindOne(id)

	// err := doc.Decode(&olive)

	// if err != nil {
	// 	apiResp.SetError(enum.ErrorCode.RecordNotFound)
	// 	return &apiResp
	// }

	// apiResp.Data = response.Olive{
	// 	Country: olive.Country,
	// 	Kind:    olive.Kind,
	// 	ID:      olive.ID.Hex(),
	// }

	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// Refresh returns a new valid token and a refresh token
func (a *Auth) Refresh() *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	var count, err = a.repository.GetOlivesCount()

	if err != nil {
		apiResp.SetError(enum.ErrorCode.RecordNotFound)
		return &apiResp
	}

	apiResp.Data = count
	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// Register creates a new user
func (a *Auth) Register() *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	var count, err = a.repository.GetOlivesCount()

	if err != nil {
		apiResp.SetError(enum.ErrorCode.RecordNotFound)
		return &apiResp
	}

	apiResp.Data = count
	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// New creates new auth api service
func New(client *mongo.Client, dbName string) *Auth {

	var uow = uow.New(client, dbName)

	var auth = Auth{}
	auth.repository = uow.GetUserRepository()
	auth.uow = uow

	return &auth
}
