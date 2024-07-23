// Package help provides utilities for handling HTTP requests.
package help

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/form"
)

// Constants for HTTP decoder.
const (
	// maxMemory is the maximum size of the request body.
	maxMemory = 4 << 20 // 4 MB
	// jsonTag is the tag name for conform parsing.
	jsonTag = "json"
)

// OptionDecoder is used to configure the decoder for custom types.
type OptionDecoder struct {
	// Func is the function used to decode custom types.
	Func form.DecodeCustomTypeFunc
	// Types is the list of types that can be decoded.
	Types []interface{}
}

// HttpDecoder is the interface that decodes HTTP requests.
type HttpDecoder interface {
	// SetMaxSize sets the maximum size of the request body.
	// The parameter limit is the maximum size in bytes.
	SetMaxSize(limit int64)

	// SetTagName sets the tag name for conform parsing.
	// The parameter tag is the name of the tag.
	SetTagName(tag string)

	// Body decodes the request body into the given interface.
	// It supports JSON and form data.
	Body(r *http.Request, i interface{}, fns ...OptionDecoder) error

	// Query decodes the query parameters into the given interface.
	Query(r *http.Request, i interface{}, fns ...OptionDecoder) error
}

// decoder is the struct that implements the HttpDecoder interface.
type decoder struct {
	maxSize int64  // The maximum size of the request body.
	tagName string // The tag name for conform parsing.
}

// NewHttpDecoder creates a new HttpDecoder with default settings.
func NewHttpDecoder() HttpDecoder {
	return &decoder{
		maxSize: maxMemory,
		tagName: jsonTag,
	}
}

// SetMaxSize sets the maximum size of the request body.
func (c *decoder) SetMaxSize(limit int64) {
	c.maxSize = limit
}

// SetTagName sets the tag name for conform parsing.
func (c *decoder) SetTagName(tag string) {
	c.tagName = tag
}

// Body decodes the request body into the given interface.
// It supports JSON and form data.
//
// Parameters:
// - r: the http.Request object.
// - i: the interface to decode into.
// - fns: the list of OptionDecoder configurations.
// Returns an error if decoding fails.
func (c *decoder) Body(r *http.Request, i interface{}, fns ...OptionDecoder) error {
	// If the request body is empty, return nil.
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	// Get the content type from the request header.
	ct := r.Header.Get("Content-Type")

	// If the content type is JSON, decode the JSON body.
	if strings.HasPrefix(ct, "application/json") || strings.HasPrefix(ct, "text/json") {
		if err := json.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	} else if strings.HasPrefix(ct, "multipart/form-data") { // If the content type is multipart/form-data, parse the multipart form.
		if err := r.ParseMultipartForm(c.maxSize); err != nil {
			return err
		}
	} else { // For other content types, parse the form.
		if err := r.ParseForm(); err != nil {
			return err
		}
	}

	// Create a new form decoder.
	d := form.NewDecoder()

	// Set the tag name for conform parsing.
	d.SetTagName(c.tagName)

	// Register custom type functions.
	for _, v := range fns {
		d.RegisterCustomTypeFunc(v.Func, v.Types...)
	}

	// Decode the form into the given interface.
	return d.Decode(i, r.Form)
}

// Query decodes the query parameters into the given interface.
//
// Parameters:
// - r: the http.Request object.
// - i: the interface to decode into.
// - fns: the list of OptionDecoder configurations.
// Returns an error if decoding fails.
func (c *decoder) Query(r *http.Request, i interface{}, fns ...OptionDecoder) error {
	// Create a new form decoder
	d := form.NewDecoder()

	// Set the tag name for conform parsing
	d.SetTagName(c.tagName)

	// Register custom type functions
	for _, v := range fns {
		d.RegisterCustomTypeFunc(v.Func, v.Types...)
	}

	// Decode the query parameters into the given interface
	return d.Decode(i, r.URL.Query())
}
