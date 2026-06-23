// CWE-915: Mass Assignment via JSON decode into struct pointer
package main
import (
	"encoding/json"
	"net/http"
)
type User struct {
	Name    string
	IsAdmin bool
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u) // IsAdmin settable from request body
}
