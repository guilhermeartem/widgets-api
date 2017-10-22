package main

import (
	"errors"
	"strconv"
	// "log"
	// "reflect"
)

//Widget class
type Widget struct {
	ID			uint 	  `json:"id";gorm:"primary_key"`
	Name		string    `json:"name";gorm:"size:255"`
	Color		string    `json:"color";gorm:"size:255"`
	Price		float64   `json:"price"`
	Inventory	int    	  `json:"inventory"`
	Melts		bool      `json:"melts"`
}

//Widgets collection
type Widgets []Widget

//CreateFromMap creats a widget from map element
func CreateFromMap(m map[string]interface{}) Widget {
	var widget Widget

	widget.Name = m["name"].(string)
	widget.Color = m["color"].(string)

	switch m["price"].(type) {
    case string:
        widget.Price, _ = strconv.ParseFloat(m["price"].(string), 64)
    default:
        widget.Price = m["price"].(float64)
    }

	switch m["inventory"].(type) {
    case string:
        widget.Inventory, _ = strconv.Atoi(m["inventory"].(string))
    default:
        widget.Inventory = int(m["inventory"].(float64))
    }

	switch m["melts"].(type) {
    case string:
        widget.Melts, _ = strconv.ParseBool(m["inventory"].(string))
    default:
        widget.Melts = m["inventory"].(bool)
    }


	return widget
}

//ValidateWidget validates a widget
func ValidateWidget(body map[string]interface{}) error {
	if val, ok := body["name"]; !ok || val == "" {
		return errors.New(".name required")
	}

	if val, ok := body["color"]; !ok || val == "" {
		return errors.New(".color required")
	}

	// log.Printf("ok: %t, val: %v, type: %v", ok, val, reflect.TypeOf(val))
	if val, ok := body["price"]; !ok || val == "" {
		return errors.New(".price required")
	}

	if val, ok := body["inventory"]; !ok || val == "" {
		return errors.New(".inventory required")
	}

	if val, ok := body["inventory"]; !ok || val == "" {
		return errors.New(".inventory required")
	}

	if val, ok := body["melts"]; !ok || val == "" {
		return errors.New(".melts required")
	}

	return nil
}
