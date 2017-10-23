package main

import (
	"log"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func redirectDocs(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/docs/", 301)
}

func main() {
	validate = validator.New()
	defer db.Close()

	router := NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/"))))
	router.HandleFunc("/docs", redirectDocs)

	handler := ValidateTokenMiddleware(UserIndex)
	router.Methods("GET").
		Path("/authenticated").
		Name("authenticated").
		Handler(handler)

	log.Fatal(http.ListenAndServe(":4000", router))
}
