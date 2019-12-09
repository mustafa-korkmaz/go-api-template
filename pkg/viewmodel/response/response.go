package response

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/message"
)

//APIResponse represents base response
type APIResponse struct {
	Result    enum.ResponseResultType `json:"result"`
	ErrorCode enum.ErrorCodeType      `json:"error_code,omitempty"`
	Message   string                  `json:"message"`
	Data      interface{}             `json:"data,omitempty"`
}

//PagedList represents a list of paginated items and total item count for the storage
type PagedList struct {
	Items        []interface{} `json:"items"`
	RecordsTotal int           `json:"records_total"`
}

//SetError sets APIResponse.Message and ErrorCode by given ErrorCode type
func (r *APIResponse) SetError(errCode enum.ErrorCodeType) {

	var msg = message.GetErrorMessage(errCode)

	if msg != nil {
		r.Message = msg.Text
		r.ErrorCode = msg.Code
	}
}
