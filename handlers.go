package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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
		fmt.Fprintln(w, "user not found")
	}
}

//WidgetIndex return list of all users
func WidgetIndex(w http.ResponseWriter, r *http.Request) {

	var widgets Widgets
	db.Find(&widgets)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(widgets); err != nil {
		panic(err)
	}
}

//WidgetShow returns one user
func WidgetShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	widgetID := vars["widgetID"]

	var widget Widget
	db.First(&widget, widgetID)

	if widget.ID != 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(widget); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "widget not found")
	}
}

//WidgetCreate creates a new Widget
func WidgetCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// var mapping interface{}
	// if err := json.Unmarshal(body, &mapping); err != nil {
	//     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//     w.WriteHeader(http.StatusUnprocessableEntity)
	//     if err := json.NewEncoder(w).Encode(err); err != nil {
	//         panic(err)
	//     }
	// 	return
	// }
	//
	// m := mapping.(map[string]interface{})
	// if err:= ValidateWidget(m); err != nil {
	// 	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	//     w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprintln(w, err)
	// 	return
	// }

	// widget := CreateFromMap(m)
	var widget Widget
	if err := json.Unmarshal(body, &widget); err != nil {
		log.Println("unmarshal error")
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		GetWidgetUnMarshalError(err)
		fmt.Fprintln(w, err)
		return
	}

	if err := widget.ValidateWidget(); err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	db.Create(&widget)
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "added!")
}

//WidgetUpdate updates a Widget
func WidgetUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	widgetID := vars["widgetID"]

	var widget Widget
	db.First(&widget, widgetID)

	if widget.ID != 0 {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &widget); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err2 := json.NewEncoder(w).Encode(err); err != nil {
				panic(err2)
			}
		}

		db.Save(&widget)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintln(w, "")
		// if err := json.NewEncoder(w).Encode(widget); err != nil {
		// 	panic(err)
		// }
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "widget not found")
	}
}
