package enum

// ResponseResultType represents enum type for response code
type ResponseResultType int

// UserStatusType represents enum type for user status
type UserStatusType byte

// ErrorCodeType represents enum type for error message
type ErrorCodeType string

// AccessLevelType represents enum type for user access level the more, the higher access level
type AccessLevelType byte

//enum list for ResponseCodeType
type responseResultTypeList struct {
	Success ResponseResultType
	Warning ResponseResultType
	Fail    ResponseResultType
}

// ResponseResult represents ResponseCodeType enums for public use
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
	None                ErrorCodeType
	AppError            ErrorCodeType
	UserNotFound        ErrorCodeType
	UserNotAuthorized   ErrorCodeType
	UserAlreadyExists   ErrorCodeType
	RecordNotFound      ErrorCodeType
	RecordAlreadyExists ErrorCodeType
	TokenCreationError  ErrorCodeType
	WeakPassword        ErrorCodeType
}

// ErrorCode represents ErrormessageType enums for public use
var ErrorCode = &errorCodeList{
	None:                "",
	AppError:            "APP_ERROR",
	UserNotFound:        "USER_NOT_FOUND",
	UserNotAuthorized:   "UNAUTHORIZED",
	RecordNotFound:      "RECORD_NOT_FOUND",
	UserAlreadyExists:   "USER_EXISTS",
	RecordAlreadyExists: "RECORD_EXISTS",
	TokenCreationError:  "TOKEN_CANNOT_CREATED",
	WeakPassword:        "WEAK_PASSWORD",
}

//enum list for ErrorMessage
type accessLevelTypeList struct {
	Standart AccessLevelType
	Author   AccessLevelType
	Admin    AccessLevelType
}

// AccessLevel represents ErrormessageType enums for public use
var AccessLevel = &accessLevelTypeList{
	Standart: 0,
	Author:   1,
	Admin:    2,
}
