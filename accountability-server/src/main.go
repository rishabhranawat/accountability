package main

import (
	"routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	
	http.Handle("/", routes.Handlers())
}