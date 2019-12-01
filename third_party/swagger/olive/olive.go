package swaggermeta

// swagger:route GET /v1/olive/count olive oliveCount
// Returns olive count stored so far
// responses:
//   200: genericResponse

// swagger:route GET /v1/olive/{id} olive oliveGet
// Returns olive details with given id.
// responses:
//   200: genericResponse
// swagger:parameters oliveGet
type oliveGetRequestWrapper struct {
	// Healthcheck GET request path.
	// in:path
	ID string `json:"id"`
}
