# Response mapper v1
This is a response mapper http server for go-helpers.

## Structure Response API 
### Error
```json
{
  "status": "status error",
  "code": 10, // code internal error
  "message": {
    "id": "message error language Indonesian",
    "en": "message error language English"
  }
}
```

### Success response message
```json
{
  "status": "Created", // status success
  "message": "data created", // response message
}
```

### Success Single Data
```json
{
  "status": "Created", // status success
  "data": {} // response data
}
```

### Success Multiple Data
```json
{
  "status": "Success", // status success
  "meta": {
    "page": 1, // current page
    "limit": 10, // current limit per page
    "total_records": 3 // total records
  },
  "data": [] // response data
}
```

## Example Usage

```go
package main

import (
	"log"
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
)

func main() {
	port := "1200"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch {
		case r.URL.Path == "/" && r.Method == "GET":
			response_mapper.RenderJSON(w, http.StatusOK, "welcome this server") // success response

		default:
			err := response_mapper.ErrRouteNotFound()
			response_mapper.RenderJSON(w, http.StatusNotFound, err) // error response
		}
	})

	log.Printf("Server is running on port %v\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

```