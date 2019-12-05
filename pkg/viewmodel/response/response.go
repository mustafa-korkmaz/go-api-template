package response

import "github.com/mustafa-korkmaz/goapitemplate/pkg/enum"

//APIResponse represents base response
type APIResponse struct {
	Code    enum.ResponseCodeType `json:"code"`
	Message string                `json:"message"`
	Data    interface{}           `json:"data"`
}

//PagedListResponse represents a list of paginated items and total item count for the storage
type PagedListResponse struct {
	Items        []interface{} `json:"items"`
	RecordsTotal int           `json:"records_total"`
}
