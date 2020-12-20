package upload

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/request"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/response"
)

// Service represents healthcheck api interface
type Service interface {
	Save(*request.Upload) *response.APIResponse
}

// Upload represents upload api service
type Upload struct {
}

// Save saves content to file share
func (u *Upload) Save(model *request.Upload) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	// var count, err = o.repository.GetOlivesCount()

	// if err != nil {
	// 	apiResp.Message = "olives cannot found"
	// 	return &apiResp
	// }

	// apiResp.Data = count
	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// New creates new Upload api service
func New() *Upload {
	var u = Upload{}
	return &u
}
