package main

import (
	"log"
	"net/http"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
)

type teamMember struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	UsernameGithub string `json:"username_github"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := teamMember{}
	decoder := help.NewHttpDecoder()
	err := decoder.Body(r, &input)
	if err != nil {
		log.Println(err)
		response_mapper.RenderJSON(w, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	response_mapper.RenderJSON(w, http.StatusOK, input)
}

func main() {
	port := "1200"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch {
		case r.Method == "GET":
			// curl --location 'http://localhost:1200'
			response_mapper.RenderJSON(w, http.StatusOK, response_mapper.MultiLanguages{
				ID: "selamat datang di server ini",
				EN: "welcome this server",
			}) // success response
			return

		case r.Method == "POST":
			/*
				curl --location 'http://localhost:1200' \
				--header 'Content-Type: application/json' \
				--data-raw '{
				    "first_name": "Adam",
				    "last_name": "nasrudin",
				    "email": "adam@example.com",
				    "username_github": "adamnasrudin03"
				}'
			*/
			ServeHTTP(w, r)
			return
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
