package v1

import (
	"net/http"
)

type ResponseError struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Err     error          `json:"-"`
	Message MultiLanguages `json:"message"`
}

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

func (e *ResponseError) Error() string {
	return e.Err.Error()
}

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

func StatusErrorMapping(code int) int {
	return statusErrorMapping[code]
}
