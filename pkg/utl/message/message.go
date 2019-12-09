package message

import "github.com/mustafa-korkmaz/goapitemplate/pkg/enum"

//Message represents a message text combined with a unuique code
type Message struct {
	Code enum.ErrorCodeType
	Text string
}

//ErrorCodeTextMap represents the code and description text for error message
var errorCodeTextMap = map[enum.ErrorCodeType]string{
	enum.ErrorCode.AppError:     "An error occured during processing request. Please try again later.",
	enum.ErrorCode.UserNotFound: "User not found.",
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
