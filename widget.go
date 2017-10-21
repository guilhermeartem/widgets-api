package main

//Widget class
type Widget struct {
	ID			uint 	  `json:"id";gorm:"primary_key"`
	Name		string    `json:"name";gorm:"size:255"`
	Color		string    `json:"color";gorm:"size:255"`
	Price		string    `json:"price";gorm:"size:255"`
	Inventory	int    	  `json:"inventory"`
	Melts		bool      `json:"melts"`
}

//Widgets collection
type Widgets []Widget
