package main

import "net/http"

//Route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes collection
type Routes []Route

var routes = Routes{
	Route{
		"UserIndex",
		"GET",
		"/users",
		UserIndex,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userID}",
		UserShow,
	},
}
