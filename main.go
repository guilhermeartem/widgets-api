package main

import (
    "log"
    "net/http"
    "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	defer db.Close()

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
