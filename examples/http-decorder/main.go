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
	Total      float64   `json:"total"` // The total number of items available.
	Date       string    `json:"date"`  // The date in string format.
	DateInTime time.Time `json:"-"`
	Limit      int       `json:"limit"`       // The maximum number of items to return.
	Offset     int       `json:"offset"`      // The number of items to skip before starting to collect the result set.
	Page       int       `json:"page"`        // The page number of results to return.
	OrderBy    string    `json:"order_by"`    // The field to order the results by.
	SortBy     string    `json:"sort_by"`     // The direction to sort the results by.
	IsNoLimit  bool      `json:"is_no_limit"` // A flag indicating whether to limit the number of items returned.
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
	input := listReq{}
	decoder := help.NewHttpDecoder()

	// Decode the query parameters into the listReq struct
	err := decoder.Query(r, &input)
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
	input := listReq{}
	decoder := help.NewHttpDecoder()

	// Decode the query parameters into the listReq struct
	err := decoder.Body(r, &input)
	if err != nil {
		APIResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	input.DateInTime, err = time.Parse(time.DateOnly, input.Date)
	if err != nil {
		APIResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	// Send the parsed boolean as the JSON response
	APIResponse(w, r, input, http.StatusOK)
}

// main starts the HTTP server and handles the request.
func main() {
	port := "1200"
	mux := http.NewServeMux()

	// Handle the "/decode-struct" route
	mux.HandleFunc("/decode-struct", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		defer help.PanicRecover("main")

		// Handle GET requests
		/*
			curl --location 'http://localhost:1200/decode-struct?page=1&limit=5&order_by=DESC&sort_by=created_at&is_no_limit=true&total=1.2345&offset=10&date=2024-07-22'
		*/
		if r.Method == "GET" {
			ServeHTTP(w, r)
			return
		}

		// Handle POST requests
		/*
			curl --location 'http://localhost:1200/decode-struct' \
			--header 'Content-Type: application/json' \
			--data '{
			    "total": 1.2345,
			    "date": "2024-07-23",
			    "limit": 5,
			    "offset": 10,
			    "page": 1,
			    "order_by": "ASC",
			    "sort_by": "created_at",
			    "is_no_limit": true
			}'
		*/
		if r.Method == "POST" {
			ServeHTTP2(w, r)
			return
		}

		// Handle other requests
		APIResponse(w, r, "Page not found", http.StatusNotFound)
	})

	// Start the HTTP server
	log.Printf("Server is running on port %v\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
