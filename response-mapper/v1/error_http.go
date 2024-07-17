package v1

import "strings"

// ResponseErrorHttp represents the structure of an error response from an HTTP request.
type ResponseErrorHttp struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Desc    MultiLanguages `json:"desc"`
	Message MultiLanguages `json:"message"`
}

// GetMessageID returns the error message in Indonesian (Bahasa Indonesia).
func (m *ResponseErrorHttp) GetMessageID() string {
	// try to get the message from the description first
	message := strings.TrimSpace(m.Desc.ID)
	if message == "" {
		// if the description is empty, try to get the message from the message itself
		message = strings.TrimSpace(m.Message.ID)
	}
	return message
}

// GetMessageEN returns the error message in English.
func (m *ResponseErrorHttp) GetMessageEN() string {
	// try to get the message from the description first
	message := strings.TrimSpace(m.Desc.EN)
	if message == "" {
		// if the description is empty, try to get the message from the message itself
		message = strings.TrimSpace(m.Message.EN)
	}
	return message
}

func ErrRouteNotFound() *ResponseError {
	return NewError(ErrNoFound, NewResponseMultiLang(
		MultiLanguages{
			ID: "Rute tidak ditemukan",
			EN: "Route not found",
		}))
}

func ErrInternalServerError() *ResponseError {
	return NewError(ErrUnknown, NewResponseMultiLang(
		MultiLanguages{
			ID: "Internal Server Error",
			EN: "Internal Server Error",
		}))
}
