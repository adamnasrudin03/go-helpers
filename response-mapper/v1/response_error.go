package v1

import (
	"net/http"
)

// ResponseError is used to represent an error response to the client.
type ResponseError struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Err     error          `json:"-"`
	Message MultiLanguages `json:"message"`
}

// NewError creates a new ResponseError from an error code and error.
//
// It sets the status, code, and message of the error based on the error code.
// If the error is already a MultiLanguages, it uses the error's message.
// If the error is not a MultiLanguages, it sets the ID and EN message to the error's message.
func NewError(code TypeError, err error) *ResponseError {
	var respErr MultiLanguages
	if errValue, isMatch := err.(*MultiLanguages); isMatch {
		if errValue != nil {
			respErr = *errValue
		} else {
			respErr = MultiLanguages{
				ID: err.Error(),
				EN: err.Error(),
			}
		}
	} else {
		respErr = MultiLanguages{
			ID: err.Error(),
			EN: err.Error(),
		}
	}
	return &ResponseError{
		Status:  StatusMapping(int(code)),
		Code:    int(code),
		Err:     err,
		Message: respErr,
	}
}

// Error returns the string representation of the error.
func (e *ResponseError) Error() string {
	return e.Err.Error()
}

// statusErrorMapping maps error codes to HTTP status codes.
var statusErrorMapping = map[int]int{
	int(ErrForbidden):    http.StatusForbidden,
	int(ErrUnauthorized): http.StatusUnauthorized,
	int(ErrDatabase):     http.StatusUnprocessableEntity,
	int(ErrFromUseCase):  http.StatusUnprocessableEntity,
	int(ErrConflict):     http.StatusConflict,
	int(ErrValidation):   http.StatusBadRequest,
	int(ErrNoFound):      http.StatusNotFound,
	int(ErrUnknown):      http.StatusInternalServerError,
}

// StatusErrorMapping returns the HTTP status code for the given error code.
//
// If the error code is not in the mapping, it returns http.StatusInternalServerError.
func StatusErrorMapping(code int) int {
	return statusErrorMapping[code]
}
