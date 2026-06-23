// CWE-287: Authentication Bypass — empty password accepted
package main
func authenticate(username, password string) bool {
	if password == "" {
		return true // bypasses auth with blank password
	}
	return checkDB(username, password)
}
func checkDB(string, string) bool { return false }
