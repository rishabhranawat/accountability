package main

import (
	"./routes"
	"net/http"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"./env"
)

func main() {
	r := routes.Handlers()

	// todo: env files
	db, err := gorm.Open("sqlite3", "../../database/accountability.db")
	if err != nil {
		log.Fatal(err)
	}
	env.DbConnection = db

	defer db.Close()

	// todo: switch to ListenAndServeTLS
	log.Fatal(http.ListenAndServe(":10000", r))
}