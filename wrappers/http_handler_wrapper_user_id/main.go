package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

type HandlerWithDBFunc func(http.ResponseWriter, *http.Request, *sql.DB)

func D(db *sql.DB, fn HandlerWithDBFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	})
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WithDbAndUserHandlerFunc func(w http.ResponseWriter, r *http.Request, db *sql.DB, user *User)

func ValidateRequest(next WithDbAndUserHandlerFunc) HandlerWithDBFunc {
	return func(w http.ResponseWriter, r *http.Request, db *sql.DB) {
		tokenString := r.Header.Get("Authorization")

		userId := parseJWT(tokenString) // get the user id from the token

		var user User
		err := db.QueryRowContext(context.Background(), "SELECT id, name FROM users WHERE id = ?", userId).
			Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err)
		}
		next(w, r, db, &user)
	}
}

func Hello(w http.ResponseWriter, r *http.Request, db *sql.DB, user *User) {
	fmt.Fprintf(w, "Hello, %s!", user.Name)
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	http.Handle("/", D(db, ValidateRequest(Hello)))
	http.ListenAndServe(":8080", nil)
}

func parseJWT(tokenString string) string {
	return tokenString
}
