package upload

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/request"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/response"
)

// Service represents healthcheck api interface
type Service interface {
	Save(*request.Upload) *response.APIResponse
	Get(fileName string) *response.APIResponse
}

// Upload represents upload api service
type Upload struct {
}

// Save saves content to file share
func (u *Upload) Save(model *request.Upload) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0750)
	}

	arr := strings.Split(model.Name, "/")
	if len(arr) > 1 {

		folder := fmt.Sprintf("data/%s", arr[0])
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			err = os.MkdirAll(folder, 0750)
			check(err)
		}
	}

	f, err := os.Create(fmt.Sprintf("data/%s", model.Name))
	check(err)

	defer f.Close()

	length, err := f.Write(model.Content)
	check(err)
	fmt.Println(length, "bytes written successfully")

	apiResp.Result = enum.ResponseResult.Success

	return &apiResp
}

// Get returns file content as byte []
func (u *Upload) Get(fileName string) *response.APIResponse {

	var apiResp = response.APIResponse{
		Result: enum.ResponseResult.Fail,
	}

	fullPath := fmt.Sprintf("data/%s", fileName)

	if _, err := os.Stat(fullPath); err == nil {
		content, err := ioutil.ReadFile(fullPath)
		check(err)
		apiResp.Data = content
	} else {
		apiResp.ErrorCode = enum.ErrorCode.RecordNotFound
	}

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
