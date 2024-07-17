package v1

import (
	"encoding/json"
	"net/http"

	help "github.com/adamnasrudin03/go-helpers"
)

// WriteJSON writes the JSON representation of v to the response writer w,
// and sets the status code of the response to statusCode. It returns an error
// if there was an error during the operation.
func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	defer help.PanicRecover("response_mapper_v1-WriteJSON")

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response header
	w.WriteHeader(statusCode)

	// Marshal the data to JSON format
	d, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// Write the JSON data to the response writer
	_, err = w.Write(d)
	return err
}

// RenderJSON renders the response based on the provided data.
func RenderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	// Check if the input data is an error
	if val, isErr := v.(error); isErr {
		if e, ok := val.(*ResponseError); ok {
			statusCode = StatusErrorMapping(e.Code)
		} else {
			val = NewError(ErrUnknown, val)
		}

		// Write the error response in JSON format
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

	// Write the response data in JSON format
	_ = WriteJSON(w, statusCode, resp)
}
