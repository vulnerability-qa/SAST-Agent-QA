// CWE-190: Integer Overflow Leading to Heap Buffer Overflow
// Vulnerability: user-controlled 'count' and 'itemSize' are multiplied to
// compute allocation size. The multiplication overflows int32 before being
// cast to int for make(), resulting in a near-zero allocation. Subsequent
// indexed writes then overflow the underlying slice, corrupting heap memory.
// Fix requires widening the arithmetic to int64 and adding overflow guards
// before the cast — not just clamping the final value.

package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func parseHeader(r io.Reader) (count int32, itemSize int32, err error) {
	buf := make([]byte, 8)
	if _, err = io.ReadFull(r, buf); err != nil {
		return
	}
	count = int32(binary.BigEndian.Uint32(buf[0:4]))
	itemSize = int32(binary.BigEndian.Uint32(buf[4:8]))
	return
}

func readItems(r io.Reader, count int32, itemSize int32) ([][]byte, error) {
	// VULNERABLE: int32 multiplication overflows before widening to int
	// e.g. count=0x10001, itemSize=0x10000 => product wraps to 0x10000
	// make() allocates only 65536 bytes; loop writes 0x10001 * 0x10000 bytes
	totalSize := int(count * itemSize) // overflow here
	if totalSize <= 0 || totalSize > 256*1024*1024 {
		return nil, fmt.Errorf("invalid total size: %d", totalSize)
	}

	buf := make([]byte, totalSize)
	items := make([][]byte, count)

	for i := int32(0); i < count; i++ {
		offset := int(i) * int(itemSize) // also overflows if i*itemSize > maxInt
		if _, err := io.ReadFull(r, buf[offset:offset+int(itemSize)]); err != nil {
			return nil, err
		}
		items[i] = buf[offset : offset+int(itemSize)]
	}
	return items, nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	count, itemSize, err := parseHeader(conn)
	if err != nil {
		return
	}
	items, err := readItems(conn, count, itemSize)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("received %d items\n", len(items))
}

func main() {
	ln, _ := net.Listen("tcp", ":9000")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
