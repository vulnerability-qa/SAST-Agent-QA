// CWE-798: Hardcoded JWT signing secret
package main
const jwtSecret = "my-super-secret-key-dont-share"
func getSecret() []byte { return []byte(jwtSecret) }
