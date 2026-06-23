// CWE-755: Unhandled panic from nil dereference reachable via user input
package main
type Config struct{ Value *string }
func getValue(c *Config) string {
	return *c.Value // panics if Value is nil (from untrusted JSON)
}
