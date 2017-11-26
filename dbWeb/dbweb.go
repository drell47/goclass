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

func (ds *databaseServer) api(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Form["last"] == nil {
		http.Error(w, fmt.Sprintf("Missing last"), http.StatusBadRequest)
		return
	}

	last := req.Form["last"][0]
	fmt.Printf("LAST NAME = %s\n", last)
	rows, err := ds.db.Query("SELECT * FROM name where last = ?;", last)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	if rows.Next() {

		var result struct {
			ID    int
			First string
			Last  string
			State bool
		}

		rows.Scan(&result.ID, &result.First, &result.Last, &result.State)
		fmt.Printf("%d, %s, %s, %t\n", result.ID, result.First, result.Last, result.State)
		j, err := json.Marshal(&result)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error %s", err), http.StatusInternalServerError)
			return
		}
		fmt.Printf("JSON: %s\n", j)
		w.Write(j)
	} else {
		fmt.Fprintf(w, "%s not found\n", last)
	}

}

// func show(db *sql.DB) {
//
// 	rows, err := db.Query("SELECT * FROM name")
// 	if err != nil {
// 		log.Fatalf("Database query error: %s", err)
// 	}
// 	defer rows.Close()
//
// 	for rows.Next() {
// 		var id int
// 		var first string
// 		var last string
// 		var state bool
// 		err = rows.Scan(&id, &first, &last, &state)
//
// 		if err != nil {
// 			log.Fatalf("Database row error: %s", err)
// 		}
// 		fmt.Printf("row type: %T\n", rows)
// 		fmt.Printf("id %d, first %s, last %s, state %t\n", id, first, last, state)
// 	}
// }

func main() {
	ds := &databaseServer{}
	var err error

	ds.db, err = sql.Open("sqlite3", "./newdb.db")
	if err != nil {
		log.Fatalf("Database open error: %s", err)
	}

	defer ds.db.Close()

	//show(db)
	http.HandleFunc("/api", ds.api)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// // start a transaction
	// t, err := db.Begin()
	// if err != nil {
	// 	log.Fatalf("Database transaction start error %s", err)
	// }
	//
	// // prepare statement to include in transation
	// st, err := t.Prepare("UPDATE name SET state = 0;")
	// if err != nil {
	// 	log.Fatalf("Database prepare error %s", err)
	// }
	//
	// // Execute
	// _, err = st.Exec()
	// if err != nil {
	// 	log.Fatalf("Database exec error %s", err)
	// }
	//
	// show(db)
	//
	// // Commit the changes - causes the change to update the database
	// t.Commit()
	//
	// show(db)

}
