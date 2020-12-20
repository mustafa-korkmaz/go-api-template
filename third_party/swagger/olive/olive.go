package swaggermeta

// swagger:route GET /v1/olives/count olives oliveCount
// Returns olive count stored so far
// responses:
//   200: genericResponse

// swagger:route GET /v1/olives/{id} olives oliveGet
// Returns olive details with given id.
// responses:
//   200: genericResponse
// swagger:parameters oliveGet
type oliveGetRequestWrapper struct {
	// Healthcheck GET request path.
	// in:path
	ID string `json:"id"`
}
