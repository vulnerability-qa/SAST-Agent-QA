// CWE-330: math/rand used for security token
package main
import (
	"crypto/rand"
	"fmt"
)
func generateToken() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%016x", b)
}
