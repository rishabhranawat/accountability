package main

import (
	"./routes"
	"net/http"
	"log"
)

func main() {
	r := routes.Handlers()
	log.Fatal(http.ListenAndServe(":10000", r))

}