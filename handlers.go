package main

import (
    "encoding/json"
	"fmt"
    "net/http"
	// "io"
	// "io/ioutil"

    "github.com/gorilla/mux"
)

//UserIndex return list of all users
func UserIndex(w http.ResponseWriter, r *http.Request) {

	var users Users
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
}

//UserShow returns one user
func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userID := vars["userID"]

	var user User
	db.First(&user, userID)

	if user.ID != 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(http.StatusOK)
	    if err := json.NewEncoder(w).Encode(user); err != nil {
	        panic(err)
	    }
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	    w.WriteHeader(http.StatusNotFound)
	    fmt.Fprintln(w, "Not Found")
	}
}
