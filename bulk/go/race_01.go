// CWE-362: Race Condition on shared counter
package main
var counter int
func increment() {
	counter++ // unsynchronized access from multiple goroutines
}
