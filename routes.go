package main

import "net/http"

//Route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	auth        bool
}

//Routes collection
type Routes []Route

var routes = Routes{
	Route{
		"UserIndex",
		"GET",
		"/users",
		UserIndex,
		true,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userID}",
		UserShow,
		true,
	},
	Route{
		"WidgetIndex",
		"GET",
		"/widgets",
		WidgetIndex,
		true,
	},
	Route{
		"WidgetShow",
		"GET",
		"/widgets/{widgetID}",
		WidgetShow,
		true,
	},
	Route{
		"WidgetCreate",
		"POST",
		"/widgets",
		WidgetCreate,
		true,
	},
	Route{
		"WidgetUpdate",
		"PUT",
		"/widgets/{widgetID}",
		WidgetUpdate,
		true,
	},
}
