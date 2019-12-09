package enum

// ResponseResultType represents enum type for response code
type ResponseResultType int

// UserStatusType represents enum type for user status
type UserStatusType byte

// ErrorCodeType represents enum type for error message
type ErrorCodeType string

//enum list for ResponseCodeType
type responseResultTypeList struct {
	Success ResponseResultType
	Warning ResponseResultType
	Fail    ResponseResultType
}

// ResponseCode represents ResponseCodeType enums for public use
var ResponseResult = &responseResultTypeList{
	Success: 1,
	Warning: 2,
	Fail:    0,
}

//enum list for UserStatus
type userStatusList struct {
	Created        UserStatusType
	EmailConfirmed UserStatusType
	Suspended      UserStatusType
	Deleted        UserStatusType
}

// UserStatus represents UserStatusType enums for public use
var UserStatus = &userStatusList{
	Created:        0,
	EmailConfirmed: 1,
	Suspended:      2,
	Deleted:        3,
}

//enum list for ErrorMessage
type errorCodeList struct {
	AppError       ErrorCodeType
	UserNotFound   ErrorCodeType
	RecordNotFound ErrorCodeType
}

// ErrorCode represents ErrormessageType enums for public use
var ErrorCode = &errorCodeList{
	AppError:       "APP_ERROR",
	UserNotFound:   "USER_NOT_FOUND",
	RecordNotFound: "RECORD_NOT_FOUND",
}
