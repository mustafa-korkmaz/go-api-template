package model

import "github.com/mustafa-korkmaz/goapitemplate/pkg/enums"

//APIResponse represents base response
type APIResponse struct {
	Code    enums.ResponseCodeType `json:"code"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
}
