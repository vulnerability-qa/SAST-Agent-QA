package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec("CREATE TABLE users (name TEXT, password TEXT)")
	db.Exec("INSERT INTO users VALUES ('alice','wonderland')")

	var username string
	fmt.Print("User: ")
	fmt.Scan(&username)

	// CWE-89: user input concatenated directly into SQL query.
	rows, _ := db.Query("SELECT * FROM users WHERE name = ?", username)
	defer rows.Close()
	for rows.Next() {
		var name, password string
		rows.Scan(&name, &password)
		fmt.Println(name, password)
	}
}
