package main

import (
	"database/sql"
	"net/http"
)

type HandlerWithDBFunc func(http.ResponseWriter, *http.Request, *sql.DB)

func D(db *sql.DB, fn HandlerWithDBFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	})
}

func Hello(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// use the db
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	http.Handle("/", D(db, Hello))
	http.ListenAndServe(":8080", nil)
}
