// CWE-89: SQL Injection via fmt.Sprintf
package main
import (
	"database/sql"
	"fmt"
)
func getUser(db *sql.DB, username string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)
	return db.Query(query)
}
