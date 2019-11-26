package swaggermeta

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck/transport"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
)

// swagger:route POST /v1/healthcheck/paginationtest healthcheck paginationtest
// Returns a sample list for pagination parameters.
// responses:
//   200: genericResponse
// swagger:parameters paginationtest
type pagedListRequestWrapper struct {
	// api pagedList request body.
	// in:body
	Body model.PagedListRequest `json:"body"`
}

// swagger:route POST /v1/healthcheck healthcheck healthcheckPost
// healthcheck checks if the api is alive.
// responses:
//   200: genericResponse
// swagger:parameters healthcheckPost
type healthcheckPostRequestWrapper struct {
	// Healthcheck POST request body.
	// in:body
	Body transport.HealthCheckReq `json:"body"`
}

// swagger:route GET /v1/healthcheck/{value} healthcheck healthcheckGet
// healthcheck checks if the api is alive.
// responses:
//   200: genericResponse
// swagger:parameters healthcheckGet
type healthcheckGetRequestWrapper struct {
	// Healthcheck GET request path.
	// in:path
	Value string `json:"value"`
}

// swagger:route GET /v2/healthcheck/{value} healthcheck healthcheckGetV2
// healthcheck checks if the api version 2 is alive.
// responses:
//   200: genericResponse
// swagger:parameters healthcheckGetV2
type healthcheckGetV2RequestWrapper struct {
	// HealthcheckV2 GET request path.
	// in:path
	Value string `json:"value"`
}
