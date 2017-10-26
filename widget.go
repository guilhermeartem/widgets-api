package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gopkg.in/go-playground/validator.v9"
)

//Widget class
type Widget struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Name      string  `validate:"required,min=1,max=255" json:"name" gorm:"size:255"`
	Color     string  `validate:"required,min=1,max=255" json:"color" gorm:"size:255"`
	Price     float64 `validate:"required" json:"price"`
	Inventory int     `validate:"required" json:"inventory"`
	Melts     *bool   `validate:"required" json:"melts"`
}

//Widgets collection
type Widgets []Widget

//MarshalJSON function for rounding price of Widget when marshalling
func (w *Widget) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        uint   `json:"id"`
		Name      string `json:"name`
		Color     string `json:"color"`
		Price     string `json:"price"`
		Inventory int    `json:"inventory"`
		Melts     *bool  `json:"melts"`
	}{
		w.ID,
		w.Name,
		w.Color,
		fmt.Sprintf("%.2f", w.Price),
		w.Inventory,
		w.Melts,
	})
}

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

//GetWidgetUnMarshalError retorna o erro de Unmarshal de forma melhor compreendida
func GetWidgetUnMarshalError(err error) error {
	types := map[string]string{
		"name":      "string",
		"color":     "string",
		"price":     "float",
		"inventory": "integer",
		"melts":     "boolean",
	}

	var returnError error

	switch err.(type) {
	case *json.UnmarshalTypeError:
		unmarshalErr := err.(*json.UnmarshalTypeError)
		returnError = fmt.Errorf("field %s should be %s, not %s", unmarshalErr.Field, types[unmarshalErr.Field], unmarshalErr.Value)
	default:
		returnError = err
	}

	return returnError
}
