package response

//Auth represents authhentication response view model
type Auth struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
