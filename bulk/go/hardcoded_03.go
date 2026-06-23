// CWE-798: Hardcoded encryption key
package main
import "crypto/aes"
var encryptionKey = []byte("1234567890123456") // hardcoded 128-bit AES key
func newCipher() (interface{}, error) {
	return aes.NewCipher(encryptionKey)
}
