package message

import "github.com/mustafa-korkmaz/goapitemplate/pkg/enum"

//Message represents a message text combined with a unuique code
type Message struct {
	Code enum.ErrorCodeType
	Text string
}

//ErrorCodeTextMap represents the code and description text for error message
var errorCodeTextMap = map[enum.ErrorCodeType]string{
	enum.ErrorCode.AppError:            "An error occured during processing request. Please try again later.",
	enum.ErrorCode.UserNotFound:        "User not found.",
	enum.ErrorCode.RecordNotFound:      "Record not found.",
	enum.ErrorCode.UserAlreadyExists:   "User already exists.",
	enum.ErrorCode.RecordAlreadyExists: "Record already exists.",
	enum.ErrorCode.WeakPassword:        "Password must be at least 8 chars and needs to include at least 1 lower, 1 upper 1 digit chars",
}

//GetErrorMessage returns Message by ErrorCode type
func GetErrorMessage(code enum.ErrorCodeType) *Message {

	var text = errorCodeTextMap[code]

	if text == "" {
		return nil
	}

	return &Message{
		Code: code,
		Text: text,
	}
}
