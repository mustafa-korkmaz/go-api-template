package request

// Login represents login viewmodel
type Login struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

// Register represents registration viewmodel
type Register struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name"`
	Email    string `json:"email" validate:"required"`
	//todo: add password validation
	Password string `json:"password" validate:"required"`
}
