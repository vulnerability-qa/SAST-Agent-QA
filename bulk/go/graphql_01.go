// CWE-89: GraphQL Injection via string interpolation
package main
import "fmt"
func buildQuery(username string) string {
	return fmt.Sprintf(`{ user(login: "%s") { name email } }`, username)
	// Attacker breaks out with: ", repos { name }}
}
