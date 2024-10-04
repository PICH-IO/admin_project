package util_error

// custom new errors
type ErrorResponse struct {
	MessageID    string `json:"message_id"`
	ErrorMessage string `json:"error_message"`
}

func NewError(message_id, error_message string) *ErrorResponse {
	return &ErrorResponse{
		MessageID:    message_id,
		ErrorMessage: error_message,
	}
}

func (e *ErrorResponse) Error() string {
	return e.ErrorMessage
}
