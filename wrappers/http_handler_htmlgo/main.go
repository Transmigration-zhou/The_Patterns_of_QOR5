package main

import (
	"context"
	"database/sql"
	. "github.com/theplant/htmlgo"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, db *sql.DB, user *User) (body HTMLComponent) {
	return Div(
		H1("Hello, " + user.Name + "!"),
	)
}

type HTMLGoHandler func(w http.ResponseWriter, r *http.Request, db *sql.DB, user *User) HTMLComponent

func HandleHTMLComponent(fn HTMLGoHandler) WithDbAndUserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, db *sql.DB, user *User) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		body := fn(w, r, db, user)
		err := Fprint(w, body, r.Context())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	http.Handle("/",
		WithDB(db,
			WithUser(
				HandleHTMLComponent(
					Hello,
				),
			),
		),
	)
	http.ListenAndServe(":8080", nil)
}

func parseJWT(tokenString string) string {
	return tokenString
}

type HandlerWithDBFunc func(http.ResponseWriter, *http.Request, *sql.DB)

func WithDB(db *sql.DB, fn HandlerWithDBFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	})
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WithDbAndUserHandlerFunc func(w http.ResponseWriter, r *http.Request, db *sql.DB, user *User)

func WithUser(next WithDbAndUserHandlerFunc) HandlerWithDBFunc {
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
