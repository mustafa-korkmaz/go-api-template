package model

//PagedListRequest represents listing conditions for paginated items
type PagedListRequest struct {
	Start               int  `json:"start"`
	Length              int  `json:"length"`
	IncludeRecordsTotal bool `json:"include_records_total"`
}

//SearchablePagedListRequest represents listing conditions with search criteria for paginated items
type SearchablePagedListRequest struct {
	PagedListRequest
	SearchCriteria interface{} `json:"search_criteria"`
}
