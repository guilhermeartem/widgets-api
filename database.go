package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//TODO use .env variables

//InitDB returns an initialized DB
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/widgets?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if db.HasTable(&User{}) == false {
		users := Users{
	        User{
				Name: "Colin",
	  			Gravatar: "http://www.gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterid",
			},
			User{
				Name: "Kyle",
	  			Gravatar: "http://www.gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterid",
			},
			User{
				Name: "Thomas",
	  			Gravatar: "http://www.gravatar.com/avatar/48009c6a27d25ef0ea03f985d1f186b0?s=64&d=monsterid",
			},
			User{
				Name: "James",
	  			Gravatar: "http://www.gravatar.com/avatar/9372f138140c8578c82bbc77b2eca602?s=64&d=monsterid",
			},
	    }

		db.CreateTable(&User{})

		for _, user := range users {
			db.Create(&user)

		}
	}

	db.AutoMigrate(&Widget{})
	return db
}

var db = InitDB()
