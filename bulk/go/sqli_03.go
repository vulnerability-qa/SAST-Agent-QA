// CWE-89: SQL Injection in GORM raw query
package main
import "gorm.io/gorm"
func findByRole(db *gorm.DB, role string) []map[string]interface{} {
	var results []map[string]interface{}
	db.Raw("SELECT * FROM users WHERE role = '" + role + "'").Scan(&results)
	return results
}
