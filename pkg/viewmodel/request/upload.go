package request

// Upload represents upload document viewmodel
type Upload struct {
	Name    string `json:"name" validate:"required"`
	Content []byte `json:"content" validate:"required"`
}
