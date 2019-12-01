package olive

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository"
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
	uow        *uow.Uow
	repository repository.OliveRepository
}

// Get returns the olive detais
func (o *Olive) Get(id string) viewmodel.APIResponse {

	var apiResp = viewmodel.APIResponse{
		Code: enum.ResponseCode.Fail,
	}

	var olive model.Olive
	var doc = o.repository.FindOne(id)

	err := doc.Decode(&olive)

	if err != nil {
		apiResp.Message = "olive id:" + id + " cannot found"
		return apiResp
	}

	apiResp.Code = enum.ResponseCode.Success

	apiResp.Data = viewmodel.Olive{
		Country: "Turkey",
		Kind:    "Gemlik",
	}

	return apiResp
}

// Count returns the total olive count
func (o *Olive) Count() viewmodel.APIResponse {

	var resp = viewmodel.APIResponse{
		Code: enum.ResponseCode.Success,
		Data: 15550,
	}

	return resp
}

// New creates new olive api service
func New(client *mongo.Client, dbName string) *Olive {

	var olive = Olive{}

	olive.uow = uow.New(client, dbName)
	olive.repository = olive.uow.OliveRepository()

	return &olive
}
