package olive

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/uow"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/response"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service represents healthcheck api interface
type Service interface {
	Get(string) *response.APIResponse
	Count() *response.APIResponse
}

//Uow represents olive service unit of work behaviour
type Uow interface {
	Save() error
	//begin transaction
}

// Repository represents user auth operations interface
type Repository interface {
	FindOneByID(id string) *mongo.SingleResult
	GetOlivesCount() (int64, error)
}

// Olive represents olive api service
type Olive struct {
	uow        Uow
	repository Repository
}

// Get returns the olive detais
func (o *Olive) Get(id string) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
		Data:   nil,
	}

	var olive model.Olive
	var doc = o.repository.FindOneByID(id)

	err := doc.Decode(&olive)

	if err != nil {
		apiResp.Message = "olive id:" + id + " cannot found"
		return &apiResp
	}

	apiResp.Data = response.Olive{
		Country: olive.Country,
		Kind:    olive.Kind,
		ID:      olive.ID.Hex(),
	}

	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// Count returns the total olive count
func (o *Olive) Count() *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	var count, err = o.repository.GetOlivesCount()

	if err != nil {
		apiResp.Message = "olives cannot found"
		return &apiResp
	}

	apiResp.Data = count
	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// New creates new olive api service
func New(client *mongo.Client, dbName string) *Olive {

	var uow = uow.New(client, dbName)

	var olive = Olive{}
	olive.repository = uow.GetOliveRepository()
	olive.uow = uow

	return &olive
}
