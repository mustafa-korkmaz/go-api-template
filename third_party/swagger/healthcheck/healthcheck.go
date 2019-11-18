package swaggermeta

import "github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck/transport"

// swagger:route GET /healthcheck healthcheck
// healthcheck checks if the api is alive.
// responses:
//   200: healthcheckResponse

// This text will appear as description of your response body.
// swagger:response healthcheckResponse
type healthcheckResponseWrapper struct {
	// in:body
	Body transport.HealthCheckResp
}

// swagger:parameters
type healthcheckRequestWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body transport.HealthCheckReq
}
