// API Docs Go api boilerplate v1
//
//     Schemes: http
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Mustafa Korkmaz <m.korkmaz@outlook.com>
//     Host: localhost:1907
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package swaggermeta

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/viewmodel/response"
)

// api generic response body.
// swagger:response genericResponse
type genericResponseWrapper struct {
	// in:body
	Body response.APIResponse
}

// api generic response body including pagedList.
type genericPagedListResponse struct {
	Result    enum.ResponseResultType `json:"result"`
	ErrorCode enum.ErrorCodeType      `json:"error_code,omitempty"`
	Message   string                  `json:"message"`
	Data      response.PagedList      `json:"data"`
}

// api generic response body including pagedList.
// swagger:response genericPagedListResponse
type genericPagedListResponseWrapper struct {
	// in:body
	Body genericPagedListResponse `json:"body"`
}
