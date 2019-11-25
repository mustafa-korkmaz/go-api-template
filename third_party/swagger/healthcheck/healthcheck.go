package swaggermeta

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck/transport"
)

// swagger:route GET /v1/healthcheck/{value} healthcheck healthcheckGet
// healthcheck checks if the api is alive.
// responses:
//   200: genericResponse

// swagger:route POST /v1/healthcheck/paginationtest healthcheck pagedListRequest
// paginationtest returns a sample list for pagination parameters.
// responses:
//   200: genericResponse

// swagger:route POST /v1/healthcheck healthcheck healthcheckPost
// healthcheck checks if the api is alive.
// responses:
//   200: genericResponse

// swagger:parameters healthcheckPost
type healthcheckPostRequestWrapper struct {
	// Healthcheck POST request body.
	// in:body
	Body transport.HealthCheckReq
}

// swagger:parameters healthcheckGet
type healthcheckGetRequestWrapper struct {
	// Healthcheck GET request path.
	// in:path
	Value string `json:"value"`
}
