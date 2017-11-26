package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type databaseServer struct {
	db *sql.DB
}

func (ds *databaseServer) post(w http.ResponseWriter, req *http.Request, last string) {
	if req.Form["first"] == nil {
		http.Error(w, fmt.Sprintf("Missing first"), http.StatusBadRequest)
		return
	}

	first := req.Form["first"][0]
	fmt.Printf("IN POST: %s, %s\n", last, first)

	// _, err := ds.db.Exec("INSERT INTO name (first, last, state) VALUES (?, ?, 1);",
	// 	first, last)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error insert %s", err), http.StatusInternalServerError)
	// 	return
	// }

	// insert
	stmt, err := ds.db.Prepare("INSERT INTO name(first, last, state) values(?,?,?)")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Prepare %s", err), http.StatusInternalServerError)
		return
	}

	res, err := stmt.Exec(first, last, 1)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Exec %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("res - %T\n", res)

	var result struct {
		Success bool
	}
	result.Success = true
	marshal(w, result)

}

func (ds *databaseServer) get(w http.ResponseWriter, req *http.Request, last string) {
	fmt.Printf("get method - LAST NAME = %s\n", last)
	rows, err := ds.db.Query("SELECT * FROM name where last = ?;", last)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var result struct {
		Success bool
		ID      int
		First   string
		Last    string
		State   bool
	}

	if rows.Next() {
		result.Success = true
		rows.Scan(&result.ID, &result.First, &result.Last, &result.State)
		fmt.Printf("%d, %s, %s, %t\n", result.ID, result.First, result.Last, result.State)

		marshal(w, result)
	} else {
		var result struct {
			Success bool
		}
		marshal(w, result)
	}

}

func marshal(w http.ResponseWriter, i interface{}) {
	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(j)
}

func (ds *databaseServer) api(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Form["last"] == nil {
		http.Error(w, fmt.Sprintf("Missing last"), http.StatusBadRequest)
		return
	}

	last := req.Form["last"][0]
	fmt.Printf("Before switch: last= %s, Method = %s\n", last, req.Method)
	switch {
	case req.Method == "GET":
		ds.get(w, req, last)
	case req.Method == "POST":
		ds.post(w, req, last)
	default:
	}

}

func main() {
	ds := &databaseServer{}
	var err error

	ds.db, err = sql.Open("sqlite3", "./newdb.db")
	if err != nil {
		log.Fatalf("Database open error: %s", err)
	}

	defer ds.db.Close()

	http.HandleFunc("/api", ds.api)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
