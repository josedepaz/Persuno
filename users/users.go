package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

// User represent information about User
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"firstname"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func findAllUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:usac2016@tcp(192.168.56.101:3306)/lmsdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, lastname, email, password FROM User")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var id int
		var name string
		var lastname string
		var email string
		var password string
		err := rows.Scan(&id, &name, &lastname, &email, &password)
		if err != nil {
			panic(err)
		}
		user := &User{ID: id, Name: name, LastName: lastname, Email: email, Password: "***"}
		users = append(users, *user)
	}
	usersList, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(usersList))
}

func createUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t User
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Sirve: %s", t)
}

func validateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t User
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Sirve: %s", t)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/"), findAllUsers)
	mux.HandleFuncC(pat.Post("/"), createUser)
	mux.HandleFuncC(pat.Post("/validate"), validateUser)

	http.ListenAndServe("localhost:8001", mux)
}
