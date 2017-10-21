package main

import "github.com/jinzhu/gorm"

//Widget class
type Widget struct {
	ID			uint 	  `json:"id";gorm:"primary_key"`
	Name		string    `json:"id";gorm:"size:255"`
	Color		string    `json:"id";gorm:"size:255"`
	Price		string    `json:"id";gorm:"size:255"`
	Inventory	int    	  `json:"id";gorm:"size:255"`
	melts		bool      `json:"id";gorm:"size:255"`
}

//Widgets collection
type Widgets []Widget
