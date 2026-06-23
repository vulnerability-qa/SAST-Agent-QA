// CWE-117: Log Injection
package main
import "log"
func logLogin(username string) {
	log.Printf("Login attempt: %s", username) // \n in username injects entries
}
