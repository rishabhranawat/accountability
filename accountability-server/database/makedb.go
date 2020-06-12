package main

import (
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"../app/src/models"
)


func main() {
	os.Remove("./accountability.db")
	db, err := gorm.Open("sqlite3", "./accountability.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.AutoMigrate(&models.User{})
}