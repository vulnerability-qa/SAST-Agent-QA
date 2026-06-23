// CWE-943: NoSQL Injection in MongoDB query via bson.M
package main
import "go.mongodb.org/mongo-driver/bson"
func buildFilter(username, password string) bson.M {
	// Attacker passes JSON object as password to bypass auth
	return bson.M{"username": username, "password": password}
}
