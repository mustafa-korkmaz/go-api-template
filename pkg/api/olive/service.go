package olive

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
	o "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository/olive"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/uow"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service represents healthcheck api interface
type Service interface {
	Get(string) viewmodel.APIResponse
	Count() viewmodel.APIResponse
}

// Olive represents olive api service
type Olive struct {
	uow        uow.Service
	repository o.Repository
}

// Get returns the olive detais
func (o *Olive) Get(id string) viewmodel.APIResponse {

	var apiResp = viewmodel.APIResponse{
		Code: enum.ResponseCode.Fail,
		Data: nil,
	}

	var olive model.Olive
	var doc = o.repository.FindOne(id)

	err := doc.Decode(&olive)

	if err != nil {
		apiResp.Message = "olive id:" + id + " cannot found"
		return apiResp
	}

	apiResp.Data = viewmodel.Olive{
		Country: olive.Country,
		Kind:    olive.Kind,
		ID:      olive.ID.Hex(),
	}

	apiResp.Code = enum.ResponseCode.Success
	return apiResp
}

// Count returns the total olive count
func (o *Olive) Count() viewmodel.APIResponse {

	var apiResp = viewmodel.APIResponse{
		Code: enum.ResponseCode.Fail,
	}

	var count, err = o.repository.GetOlivesCount()

	if err != nil {
		apiResp.Message = "olives cannot found"
		return apiResp
	}

	apiResp.Data = count

	apiResp.Code = enum.ResponseCode.Success
	return apiResp
}

// New creates new olive api service
func New(client *mongo.Client, dbName string) *Olive {

	var olive = Olive{}

	olive.uow = uow.New(client, dbName)
	olive.repository = olive.uow.OliveRepository()

	return &olive
}
