package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	os.Remove("./accountability.db")
	db, err := sql.Open("sqlite3", "./accountability.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := "create table task (id integer not null primary key, name text, description text);"

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}