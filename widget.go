package main

import (
	"errors"
	// "strconv"
	"log"
	// "reflect"
	"gopkg.in/go-playground/validator.v9"
)

//Widget class
type Widget struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Name      string  `validate:"required,min=1,max=255" json:"name" gorm:"size:255"`
	Color     string  `validate:"required,min=1,max=255" json:"color" gorm:"size:255"`
	Price     float64 `validate:"required" json:"price,string,omitempty"`
	Inventory int     `validate:"required" json:"inventory"`
	Melts     *bool   `validate:"required" json:"melts"`
}

//Widgets collection
type Widgets []Widget

//ValidateWidget validates a widget
func (widget *Widget) ValidateWidget() error {
	log.Println(widget)
	err := validate.Struct(widget)
	if err != nil {
		log.Println("aqui")
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			return err
		}
		field := err.(validator.ValidationErrors)[0].Field()
		switch field {
		case "Name":
			return errors.New(".name required")
		case "Color":
			return errors.New(".color required")
		case "Price":
			return errors.New(".price required")
		case "Inventory":
			return errors.New(".inventory required")
		case "Melts":
			return errors.New(".melts required")
		default:
			log.Println(err)
		}
		return err
	}
	return nil
}

func GetWidgetUnMarshalError(err error) {
	log.Println(err)
}

// //CreateFromMap creats a widget from map element
// func CreateFromMap(m map[string]interface{}) Widget {
// 	var widget Widget

// 	widget.Name = m["name"].(string)
// 	widget.Color = m["color"].(string)

// 	switch m["price"].(type) {
// 	case string:
// 		widget.Price, _ = strconv.ParseFloat(m["price"].(string), 64)
// 	default:
// 		widget.Price = m["price"].(float64)
// 	}

// 	switch m["inventory"].(type) {
// 	case string:
// 		widget.Inventory, _ = strconv.Atoi(m["inventory"].(string))
// 	default:
// 		widget.Inventory = int(m["inventory"].(float64))
// 	}

// 	switch m["melts"].(type) {
// 	case string:
// 		widget.Melts, _ = strconv.ParseBool(m["inventory"].(string))
// 	default:
// 		widget.Melts = m["inventory"].(bool)
// 	}

// 	return widget
// }

// //ValidateWidget validates a widget
// func ValidateWidget(body map[string]interface{}) error {
// 	if val, ok := body["name"]; !ok || val == "" {
// 		return errors.New(".name required")
// 	}

// 	if val, ok := body["color"]; !ok || val == "" {
// 		return errors.New(".color required")
// 	}

// 	// log.Printf("ok: %t, val: %v, type: %v", ok, val, reflect.TypeOf(val))
// 	if val, ok := body["price"]; !ok || val == "" {
// 		return errors.New(".price required")
// 	}

// 	if val, ok := body["inventory"]; !ok || val == "" {
// 		return errors.New(".inventory required")
// 	}

// 	if val, ok := body["inventory"]; !ok || val == "" {
// 		return errors.New(".inventory required")
// 	}

// 	if val, ok := body["melts"]; !ok || val == "" {
// 		return errors.New(".melts required")
// 	}

// 	return nil
// }
