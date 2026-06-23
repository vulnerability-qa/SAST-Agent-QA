// CWE-190: Integer overflow in buffer allocation
package main
func allocate(count int32, itemSize int32) []byte {
	total := count * itemSize // overflows if large values supplied
	return make([]byte, total)
}
