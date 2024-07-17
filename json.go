package help

import "encoding/json"

// SafeJsonMarshal is a function that safely marshal a data to json.
// If the Marshal process is failed, it will return the original Marshal result
// and the error. This function is useful when you want to ensure that the data
// is marshaled to json without any errors.
//
// Parameters:
// - data: the data to be marshaled to json.
//
// Returns:
// - []byte: the marshaled json data.
// - error: the error if the Marshal process is failed, otherwise nil.
func SafeJsonMarshal(data interface{}) ([]byte, error) {
	// Marshal the data to json
	resp, err := json.Marshal(data)
	if err != nil {
		// If the Marshal process is failed, return the original Marshal result
		// and the error.
		return resp, err
	}

	// Unmarshal the json data to a temporary data
	var tempData interface{}
	err = json.Unmarshal(resp, &tempData)
	if err != nil {
		// If the Unmarshal process is failed, return the original Marshal result
		// and the error.
		return json.Marshal(data)
	}

	// Return the json data and nil error
	return resp, nil
}

// JsonToStruct is a function that unmarshal a json string to a struct.
//
// Parameters:
// - params: the json string to be unmarshaled to the struct.
// - data: the struct to be unmarshaled to.
//
// Returns:
// - error: the error if the Unmarshal process is failed, otherwise nil.
func JsonToStruct(params string, data interface{}) error {
	// Unmarshal the json string to the struct
	err := json.Unmarshal([]byte(params), data)
	if err != nil {
		// Return the error
		return err
	}

	// Return nil error
	return nil
}

// JsonToString is a function that marshal a struct to json string.
//
// Parameters:
// - data: the struct to be marshaled to json.
//
// Returns:
// - string: the marshaled json string.
// - error: the error if the Marshal process is failed, otherwise nil.
func JsonToString(data interface{}) (string, error) {
	// Marshal the struct to json
	resp, err := SafeJsonMarshal(data)
	if err != nil {
		// Return the error
		return "", err
	}

	// Return the json string
	return string(resp), nil
}
