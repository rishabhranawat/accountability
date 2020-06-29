package main

import (
	"log"
	"net/http"

	"./env"
	"./routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"http://localhost:4200"},                 // All origins
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}, // Allowing only get, just an example
		AllowCredentials: true,
	})

	// todo: switch to ListenAndServeTLS
	log.Fatal(http.ListenAndServe(":10000", c.Handler(r)))
}
