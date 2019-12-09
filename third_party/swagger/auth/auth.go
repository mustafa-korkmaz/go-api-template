package swaggermeta

// swagger:route GET /v1/refresh auth refreshToken
// Returns new valid access_token and refresh_token
// responses:
//   200: genericResponse

type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

// swagger:route POST /v1/register auth register
// Creates new user
// responses:
//   200: genericResponse
// swagger:parameters register
type registerRequestWrapper struct {
	// registerRequest POST request bofy.
	// in:body
	Body registerRequest `json:"body"`
}

type loginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

// swagger:route POST /v1/login auth login
// Authenticates user and returns valid access_token and refresh_token
// responses:
//   200: genericResponse
// swagger:parameters login
type loginRequestWrapper struct {
	// LoginRequest POST request body.
	// in:body
	Body loginRequest `json:"body"`
}
