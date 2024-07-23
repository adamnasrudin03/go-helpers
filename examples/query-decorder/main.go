package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
)

// listReq represents the list request entity with struct tags.
// The struct tags are used to determine which query parameters belong to which fields.
// The struct tags are in the format "query:<key>".
type listReq struct {
	Limit      int       `query:"limit"`         // The maximum number of items to return.
	Offset     int       `query:"offset"`        // The number of items to skip before starting to collect the result set.
	Page       int       `query:"page"`          // The page number of results to return.
	OrderBy    string    `query:"order_by"`      // The field to order the results by.
	SortBy     string    `query:"sort_by"`       // The direction to sort the results by.
	IsNoLimit  bool      `query:"is_no_limit"`   // A flag indicating whether to limit the number of items returned.
	Total      float64   `query:"total"`         // The total number of items available.
	Date       string    `query:"date" json:"-"` // The date in string format.
	DateInTime time.Time `json:"date"`           // The date in time.Time format.
}

// APIResponse sends a JSON response with the given data and status code.
// If there is an error marshaling the data, it sends an error response with the corresponding status code.
//
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request object.
// - data: The data to be encoded as JSON.
// - code: The HTTP status code of the response.
func APIResponse(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		APIResponse(w, r, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(jsonBytes)
}

// ServeHTTP parses the query parameters into a listReq struct.
// It uses the struct tags to determine which query parameters belong to which fields.
//
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request object.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of the listReq struct
	input := listReq{}
	query := r.URL.Query()

	// Decode the query parameters into the listReq struct
	err := help.NewQueryDecoder(query).Decode(&input)
	if err != nil {
		APIResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	input.DateInTime, err = time.Parse(time.DateOnly, input.Date)
	if err != nil {
		APIResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	// Send the decoded listReq struct as the JSON response
	APIResponse(w, r, input, http.StatusOK)
}

// ServeHTTP2 parses the query parameter "is_no_limit" into a boolean.
// It uses the query key to determine which field to decode into.
//
// Parameters:
// - w: The http.ResponseWriter to write the response to.
// - r: The http.Request object.
func ServeHTTP2(w http.ResponseWriter, r *http.Request) {
	var isNoLimit bool
	query := r.URL.Query()

	// Decode the "is_no_limit" query parameter into a boolean variable
	err := help.NewQueryDecoder(query).DecodeField("is_no_limit", &isNoLimit)
	if err != nil {
		APIResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	// Send the parsed boolean as the JSON response
	APIResponse(w, r, isNoLimit, http.StatusOK)
}

// main starts the HTTP server and handles the request.
func main() {
	port := "1200"
	mux := http.NewServeMux()

	// Handle the "/decode-field" route
	mux.HandleFunc("/decode-field", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		defer help.PanicRecover("main")

		if r.Method == "GET" {
			// http://localhost:1200/decode-field?is_no_limit=true
			ServeHTTP2(w, r)
			return
		}
		APIResponse(w, r, "Page not found", http.StatusNotFound)
	})

	// Handle the "/decode-struct" route
	mux.HandleFunc("/decode-struct", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		defer help.PanicRecover("main")

		if r.Method == "GET" {
			// http://localhost:1200/decode-struct?page=1&limit=5&order_by=DESC&sort_by=created_at&is_no_limit=true&total=1.2345&offset=10&date=2024-07-22
			ServeHTTP(w, r)
			return
		}
		APIResponse(w, r, "Page not found", http.StatusNotFound)
	})

	// Start the HTTP server
	log.Printf("Server is running on port %v\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
