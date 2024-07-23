package help

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/form"
)

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

// NewHttpDecoder creates a new HttpDecoder.
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
// r: the http.Request object.
// i: the interface to decode into.
// fns: the list of OptionDecoder configurations.
// Returns an error if decoding fails.
func (c *decoder) Body(r *http.Request, i interface{}, fns ...OptionDecoder) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") || strings.HasPrefix(ct, "text/json") {
		if err := json.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	} else if strings.HasPrefix(ct, "multipart/form-data") {
		if err := r.ParseMultipartForm(c.maxSize); err != nil {
			return err
		}
	} else {
		if err := r.ParseForm(); err != nil {
			return err
		}
	}

	d := form.NewDecoder()
	d.SetTagName(c.tagName)
	for _, v := range fns {
		d.RegisterCustomTypeFunc(v.Func, v.Types...)
	}

	return d.Decode(i, r.Form)
}

// Query decodes the query parameters into the given interface.
//
// r: the http.Request object.
// i: the interface to decode into.
// fns: the list of OptionDecoder configurations.
// Returns an error if decoding fails.
func (c *decoder) Query(r *http.Request, i interface{}, fns ...OptionDecoder) error {
	d := form.NewDecoder()
	d.SetTagName(c.tagName)
	for _, v := range fns {
		d.RegisterCustomTypeFunc(v.Func, v.Types...)
	}

	return d.Decode(i, r.URL.Query())
}
