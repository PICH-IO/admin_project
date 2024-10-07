package util_error

import "fmt"

type ErrorResponse struct {
	MessageID    string `json:"message_id"`
	ErrorMessage string `json:"error_message"`
}

func NewError(messageID, errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		MessageID:    messageID,
		ErrorMessage: errorMessage,
	}
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error %s: %s", e.MessageID, e.ErrorMessage)
}
