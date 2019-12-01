package olive

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel"
)

// Service represents healthcheck api interface
type Service interface {
	Get(string) viewmodel.APIResponse
	Count() viewmodel.APIResponse
}

// Olive represents olive api service
type Olive struct{}

// Get returns the olive detais
func (o *Olive) Get(id string) viewmodel.APIResponse {

	var resp = viewmodel.APIResponse{
		Code: enum.ResponseCode.Success,
		Data: viewmodel.Olive{
			Country: "Turkey",
			Kind:    "Gemlik",
		},
	}

	return resp
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
func New() *Olive {
	return new(Olive)
}
