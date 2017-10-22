package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

//PopulateUsers populates the users table with initial data
func PopulateUsers(db *gorm.DB) {
	users := Users{
		User{
			Name:     "Colin",
			Gravatar: "http://www.gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterid",
		},
		User{
			Name:     "Kyle",
			Gravatar: "http://www.gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterid",
		},
		User{
			Name:     "Thomas",
			Gravatar: "http://www.gravatar.com/avatar/48009c6a27d25ef0ea03f985d1f186b0?s=64&d=monsterid",
		},
		User{
			Name:     "James",
			Gravatar: "http://www.gravatar.com/avatar/9372f138140c8578c82bbc77b2eca602?s=64&d=monsterid",
		},
	}

	for _, user := range users {
		db.Create(&user)

	}
}

//env loads a single env variable, and if it doesn't exists, returns the default value
func env(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		val = def
	}

	return val
}

//InitDB returns an initialized DB
//TODO use .env variables
func InitDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found!")
	}

	dburi := env("DB_USERNAME", "user") + ":" + env("DB_PASSWORD", "password") +
		"@tcp(" + env("DB_HOST", "localhost") + ":" + env("DB_PORT", "3306") +
		")/" + env("DB_DATABASE", "database") + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dburi)
	if err != nil {
		panic(err)
	}

	if env("DB_AUTO_MIGRATE", "0") == "1" {
		if db.HasTable(&User{}) == false {
			if env("DB_AUTO_POPULATE", "0") == "1" {
				db.CreateTable(&User{})
				PopulateUsers(db)
			}

		}

		db.AutoMigrate(&User{}, &Widget{})
	}
	return db
}

var db = InitDB()
