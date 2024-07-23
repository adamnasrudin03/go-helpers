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
// It writes the response data in JSON format with the specified status code.
// If the input data is an error, it sets the status code according to the error code.
func RenderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	// Create the response structure based on the input data
	resp := RenderStruct(statusCode, v)

	// Check if the input data is an error
	if val, isErr := v.(error); isErr {
		// If the input data is an error, set the status code accordingly
		if e, ok := val.(*ResponseError); ok {
			statusCode = StatusErrorMapping(e.Code)
		}
	}

	// Write the response data in JSON format with the specified status code
	_ = WriteJSON(w, statusCode, resp)
}

// RenderStruct renders the response based on the provided data.
// It creates a ResponseDefault structure based on the input data type.
func RenderStruct(statusCode int, v interface{}) interface{} {
	// Check if the input data is an error
	if val, isErr := v.(error); isErr {
		// If the input data is an error, create a ResponseError structure
		e, ok := val.(*ResponseError)
		if !ok {
			e = NewError(ErrUnknown, val)
		}
		return e
	}

	var resp ResponseDefault
	switch data := v.(type) {
	case *Pagination:
		// If the input data is a Pagination structure, create a ResponseDefault structure with the pagination data
		paginate := data
		resp = ResponseDefault{
			Status: StatusMapping(statusCode),
			Meta:   paginate.Meta,
			Data:   paginate.Data,
		}
	case MultiLanguages:
		// If the input data is a MultiLanguages structure, create a ResponseDefault structure with the message data
		resp = ResponseDefault{
			Status:  StatusMapping(statusCode),
			Message: data,
		}
	case string:
		// If the input data is a string, create a ResponseDefault structure with the message data
		resp = ResponseDefault{
			Status:  StatusMapping(statusCode),
			Message: data,
		}
	default:
		// If the input data is of any other type, create a ResponseDefault structure with the data
		resp = ResponseDefault{
			Status: StatusMapping(statusCode),
			Data:   data,
		}
	}

	return resp
}
