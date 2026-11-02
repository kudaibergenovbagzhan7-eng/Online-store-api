package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connection := "user=postgres password=Aa1234 dbname=online_store_apiDB host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
}
func CloseDB() {
	var err error
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	InitDB()
	defer CloseDB()

	server := http.Server{
		Addr: "localhost:5050",
	}
	fmt.Println("Server application is starting", server.Addr)
	server.ListenAndServe()
}
