package enums

// ResponseCodeType represents enum type for response code
type ResponseCodeType int

//enum list for ResponseCodeType
type list struct {
	Success ResponseCodeType
	Warning ResponseCodeType
	Fail    ResponseCodeType
}

// ResponseCode for public use
var ResponseCode = &list{
	Success: 1,
	Warning: 2,
	Fail:    0,
}
