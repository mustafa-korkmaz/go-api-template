package request

//PagedListRequest represents listing conditions for paginated items
type PagedList struct {
	Start               int  `json:"start"`
	Length              int  `json:"length"`
	IncludeRecordsTotal bool `json:"include_records_total"`
}

//SearchablePagedListRequest represents listing conditions with search criteria for paginated items
type SearchablePagedList struct {
	PagedList
	SearchCriteria interface{} `json:"search_criteria"`
}
