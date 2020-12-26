package upload

import (
	"fmt"
	"os"
	"strings"

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

	arr := strings.Split(model.Name, "/")
	if len(arr) > 1 {
		err := os.Mkdir(fmt.Sprintf("data/%s", arr[0]), 0750)

		check(err)
	}

	err := os.Mkdir("data", 0750)
	f, err := os.Create(fmt.Sprintf("data/%s", model.Name))
	check(err)

	defer f.Close()

	length, err := f.Write(model.Content)
	check(err)
	fmt.Println(length, "bytes written successfully")

	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// New creates new Upload api service
func New() *Upload {
	var u = Upload{}
	return &u
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
