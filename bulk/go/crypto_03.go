// CWE-330: math/rand used for security token
package main
import (
	"fmt"
	"math/rand"
)
func generateToken() string {
	return fmt.Sprintf("%016x", rand.Int63()) // not cryptographically secure
}
