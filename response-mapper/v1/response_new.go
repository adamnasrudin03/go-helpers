package v1

import (
	"encoding/json"
	"net/http"

	go_helpers_error "github.com/adamnasrudin03/go-helpers/error"
)

// WriteJSON writes the JSON representation of v to the response writer w,
// and sets the status code of the response to statusCode. It returns an error
// if there was an error during the operation.
func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	defer go_helpers_error.PanicRecover("response_mapper_v1-WriteJSON")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	d, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}

func RenderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	if val, isErr := v.(error); isErr {
		statusCode = StatusCodeMapping(statusCode, val)
		_ = WriteJSON(w, statusCode, val)
		return
	}

	var resp ResponseDefault
	switch data := v.(type) {
	case *Pagination:
		paginate := data
		resp = ResponseDefault{
			Status: StatusMapping(statusCode),
			Meta:   paginate.Meta,
			Data:   paginate.Data,
		}
	case string:
		resp = ResponseDefault{
			Status:  StatusMapping(statusCode),
			Message: v.(string),
		}
	default:
		resp = ResponseDefault{
			Status: StatusMapping(statusCode),
			Data:   data,
		}
	}
	_ = WriteJSON(w, statusCode, resp)
}

func StatusCodeMapping(statusCode int, v interface{}) int {
	if e, ok := v.(*ResponseError); ok {
		statusCode = StatusErrorMapping(e.Code)
	}
	return statusCode
}