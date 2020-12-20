package swaggermeta

type uploadRequest struct {
	Name    string `json:"name" validate:"required"`
	Content []byte `json:"content" validate:"required"`
}

// swagger:route POST /v1/uploads uploads uploadPost
// Saves the content
// responses:
//   200: genericResponse
// swagger:parameters uploadPost
type uploadPostRequestWrapper struct {
	// uploadRequest POST request body.
	// in:body
	Body uploadRequest `json:"body"`
}
