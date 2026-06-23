// CWE-89: SQL Injection via string concatenation
package main
import "database/sql"
func searchProducts(db *sql.DB, keyword string) (*sql.Rows, error) {
	return db.Query("SELECT * FROM products WHERE name LIKE '%" + keyword + "%'")
}
