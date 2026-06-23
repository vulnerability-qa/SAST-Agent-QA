// CWE-798: Hardcoded database password
package main
const dbPassword = "P@ssw0rd!Secret"
func getDSN() string {
	return "postgres://admin:" + dbPassword + "@db.internal:5432/prod"
}
