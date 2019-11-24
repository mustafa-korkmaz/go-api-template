package model

//PagedList represents a list of items with pagination
type PagedList struct {
	Items        []interface{} `json:"items"`
	RecordsTotal int           `json:"records_total"`
}
