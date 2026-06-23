// CWE-362: Race Condition on balance check-then-act
package main
var balance = 1000
func withdraw(amount int) bool {
	if balance >= amount { // check
		balance -= amount // update — no mutex
		return true
	}
	return false
}
