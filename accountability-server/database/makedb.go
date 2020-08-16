package main

import (
	"log"

	"../app/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// os.Remove("./accountability.db")
	db, err := gorm.Open("sqlite3", "./accountability.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Task{})
	// db.AutoMigrate(&models.TaskUpdate{})
	// db.AutoMigrate(&models.Tracker{})

	db.AutoMigrate(&models.TaskComment{})
}
