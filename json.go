package help

import "encoding/json"

// SafeJsonMarshal is a function that safely marshal a data to json.
// if the Marshal process is failed, it will return the original Marshal result
// and the error. This function is useful when you want to ensure that the data
// is marshaled to json without any errors.
func SafeJsonMarshal(data interface{}) ([]byte, error) {
	resp, err := json.Marshal(data)
	if err != nil {
		return resp, err
	}
	var tempData interface{}
	err = json.Unmarshal(resp, &tempData)
	if err != nil {
		return json.Marshal(data)
	}
	return resp, nil
}
