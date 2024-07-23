package help

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

// QueryDecoder is an interface for decoding query parameters into a struct.
type QueryDecoder interface {
	// Decode decodes the query parameters into the target struct.
	Decode(target interface{}) error
	// DecodeField decodes the query parameter with the given key into the target field.
	DecodeField(key string, target interface{}) error
}

// queryDecoder is a struct that implements the QueryDecoder interface.
type queryDecoder struct {
	values url.Values
}

// New creates a new queryDecoder instance with the given values.
func NewQueryDecoder(values url.Values) QueryDecoder {
	// creates a new queryDecoder instance
	return &queryDecoder{
		values: values,
	}
}

// DecodeField decodes the query parameter with the given key into the target field.
func (q *queryDecoder) DecodeField(key string, target interface{}) error {
	if v, ok := q.values[key]; ok {
		// parse and set value for the field
		return parseAndSetValue(v[0], target)
	}
	return nil
}

// Decode decodes the query parameters into the target struct.
func (q *queryDecoder) Decode(target interface{}) error {
	rv := reflect.ValueOf(target).Elem()
	for i := 0; i < rv.NumField(); i++ {
		key := rv.Type().Field(i).Tag.Get("query")
		if key == "" {
			continue
		}
		if err := q.DecodeField(key, rv.Field(i).Addr().Interface()); err != nil {
			return err
		}
	}
	return nil
}

// parseAndSetValue parses the string and sets the value of the reflect value.
func parseAndSetValue(s string, target interface{}) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// switch based on the target type to parse and set value
	switch rv.Kind() {
	case reflect.String:
		rv.SetString(s)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		rv.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, rv.Type().Bits())
		if err != nil {
			return err
		}
		rv.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(s, 10, rv.Type().Bits())
		if err != nil {
			return err
		}
		rv.SetUint(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(s, rv.Type().Bits())
		if err != nil {
			return err
		}
		rv.SetFloat(n)
	default:
		return errors.New("unsupported type")
	}
	return nil
}
