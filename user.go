package main

//User class
type User struct {
	ID			uint 	  `json:"id";gorm:"primary_key"`
	Name		string    `json:"name";gorm:"size:255"`
	Gravatar	string    `json:"gravatar";gorm:"size:255"`
}

//Users collection
type Users []User
