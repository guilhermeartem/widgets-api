package main

import (
    // "encoding/json"
    "fmt"
    "net/http"
	// "io"
	// "io/ioutil"

    // "github.com/gorilla/mux"
)

//Index shows the index of the page
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
