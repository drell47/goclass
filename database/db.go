package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func show(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM name")
	if err != nil {
		log.Fatalf("Database query error: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var first string
		var last string
		var state bool
		err = rows.Scan(&id, &first, &last, &state)

		if err != nil {
			log.Fatalf("Database row error: %s", err)
		}
		fmt.Printf("row type: %T\n", rows)
		fmt.Printf("id %d, first %s, last %s, state %t\n", id, first, last, state)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./newdb.db")
	if err != nil {
		log.Fatalf("Database open error: %s", err)
	}
	fmt.Printf("Type for db = %T\n", db)

	defer db.Close()

	show(db)

	// start a transaction
	t, err := db.Begin()
	if err != nil {
		log.Fatalf("Database transaction start error %s", err)
	}

	// prepare statement to include in transation
	st, err := t.Prepare("UPDATE name SET state = 0;")
	if err != nil {
		log.Fatalf("Database prepare error %s", err)
	}

	// Execute
	_, err = st.Exec()
	if err != nil {
		log.Fatalf("Database exec error %s", err)
	}

	show(db)

	// Commit the changes - causes the change to update the database
	t.Commit()

	show(db)

}
