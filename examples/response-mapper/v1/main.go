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
			response_mapper.RenderJSON(w, http.StatusOK, response_mapper.MultiLanguages{
				ID: "selamat datang di server ini",
				EN: "welcome this server",
			}) // success response

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
