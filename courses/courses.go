package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
    "database/sql"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

// Course represent information about Course
type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"firstname"`
	Description string `json:"description"`
}

func findAllCourses(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:usac2016@tcp(mysql:3306)/lmsdb")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name, description FROM Course")
	if err != nil {
		panic(err)
	}
	var courses []Course
	for rows.Next() {
		var id int
		var name string
		var description string

		err := rows.Scan(&id, &name, &description)
		if err != nil {
			panic(err)
		}
		course := &Course{ID: id, Name: name, Description: description}
		courses = append(courses, *course)
	}
	coursesList, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(coursesList))
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/"), findAllCourses)

	http.ListenAndServe("localhost:8001", mux)
}
