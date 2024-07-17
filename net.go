package help

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	IsStatusSuccess = map[int]bool{
		http.StatusOK:      true,
		http.StatusCreated: true,
	}
)

// StreamToString converts an io.Reader to a string.
//
// Parameters:
// - stream: The io.Reader to be converted.
//
// Returns:
// - A string representation of the io.Reader, or an empty string if the input is nil.
func StreamToString(stream io.Reader) string {
	// If the input stream is nil, return an empty string.
	if stream == nil {
		return ""
	}

	// Create a new Buffer to read from the stream.
	buf := new(bytes.Buffer)

	// Read the stream into the Buffer.
	buf.ReadFrom(stream)

	// Return the string representation of the Buffer.
	return buf.String()
}

// StreamToByte converts an io.Reader to a byte slice.
//
// Parameters:
// - stream: The io.Reader to be converted.
//
// Returns:
// - A byte slice representation of the io.Reader, or an empty byte slice if the input is nil.
func StreamToByte(stream io.Reader) []byte {
	// If the input stream is nil, return an empty byte slice.
	if stream == nil {
		return []byte{}
	}

	// Create a new Buffer to read from the stream.
	buf := new(bytes.Buffer)

	// Read the stream into the Buffer.
	_, err := buf.ReadFrom(stream)
	if err != nil {
		log.Printf("error reading from stream: %v", err)
		return []byte{}
	}

	// Return the byte slice representation of the Buffer.
	return buf.Bytes()
}

// GetHTTPRequestJSON sends an HTTP request with the given method, URL, body, and timeout, and returns the response as a byte slice.
// The function takes an optional set of headers to include with the request.
//
// Parameters:
// - ctx: The context to use for the request.
// - method: The HTTP method to use (e.g. "GET", "POST", etc.).
// - url: The URL to send the request to.
// - body: The body of the request to send.
// - customTimeOut: The timeout to use for the request, in seconds.
// - headers: An optional set of headers to include with the request.
//
// Returns:
// - res: The response from the server as a byte slice.
// - statusCode: The HTTP status code of the response.
// - err: An error if the request fails.
func GetHTTPRequestJSON(ctx context.Context, method string, url string, body io.Reader, customTimeOut int, headers ...map[string]string) (res []byte, statusCode int, err error) {
	defer PanicRecover("net-GetHTTPRequestJSON")

	// Create an HTTP request with the given method, URL, and body.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Iterate over the optional set of headers and add them to the request.
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	// Create an HTTP client with the given timeout.
	client := &http.Client{Timeout: time.Duration(customTimeOut) * time.Second}

	// Send the request and get the response.
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Read the response body into a byte slice.
	resp := StreamToByte(r.Body)

	// Close the response body and log the request and response.
	defer func() {
		r.Body.Close()
		// tags := map[string]interface{}{
		// 	"http.headers":    req.Header,
		// 	"http.method":     req.Method,
		// 	"http.url":        req.URL.String(),
		// 	"response.status": r.Status,
		// 	"response.body":   string(resp),
		// }
		// log.Println(tags)
	}()

	// Return the response body, status code, and any error.
	return resp, r.StatusCode, nil
}

// GetHTTPRequestSkipVerify sends an HTTP request with the given method, URL, body, and timeout,
// and returns the response as a byte slice.
// The function takes an optional set of headers to include with the request.
//
// Skips SSL certificate verification.
//
// Parameters:
// - ctx: The context to use for the request.
// - method: The HTTP method to use (e.g. "GET", "POST", etc.).
// - url: The URL to send the request to.
// - body: The body of the request to send.
// - customTimeOut: The timeout to use for the request, in seconds.
// - headers: An optional set of headers to include with the request.
//
// Returns:
// - res: The response from the server as a byte slice.
// - statusCode: The HTTP status code of the response.
// - err: An error if the request fails.
func GetHTTPRequestSkipVerify(ctx context.Context, method string, url string, body io.Reader, customTimeOut int, headers ...map[string]string) (res []byte, statusCode int, err error) {
	defer PanicRecover("net-GetHTTPRequestSkipVerify")

	// Create an HTTP request with the given method, URL, and body.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Iterate over the optional set of headers and add them to the request.
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	// Create an HTTP client with the given timeout.
	client := &http.Client{Timeout: time.Duration(customTimeOut) * time.Second}

	// Skip SSL certificate verification.
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Send the request and get the response.
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Read the response body into a byte slice.
	resp := StreamToByte(r.Body)

	// Close the response body and log the request and response.
	defer func() {
		r.Body.Close()
		// tags := map[string]interface{}{
		// 	"http.headers":    req.Header,
		// 	"http.method":     req.Method,
		// 	"http.url":        req.URL.String(),
		// 	"response.status": r.Status,
		// 	"response.body":   string(resp),
		// }
		// log.Println(tags)
	}()

	// Return the response body, status code, and any error.
	return resp, r.StatusCode, nil
}
