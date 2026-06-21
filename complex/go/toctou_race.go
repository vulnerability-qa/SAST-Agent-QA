// CWE-367: TOCTOU Race Condition
// Vulnerability: file permission is checked (T) then the file is opened (U)
// in a separate goroutine with no synchronization. An attacker who can
// swap the file between the check and the open can bypass the permission guard.
// Fix requires opening the file first and checking permissions on the descriptor,
// not on the path — a structural change, not a simple patch.

package main

import (
	"fmt"
	"os"
	"sync"
)

var mu sync.Mutex

func isFilePermitted(path string, requiredMode os.FileMode) bool {
	// CHECK (T): reads file metadata at this moment in time
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&requiredMode == requiredMode
}

func processFile(path string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// TOCTOU: check happens here...
	if !isFilePermitted(path, 0o444) {
		results <- fmt.Sprintf("denied: %s", path)
		return
	}

	// ...USE happens here, in a separate goroutine call stack.
	// An attacker process can swap path to a symlink pointing at /etc/shadow
	// between isFilePermitted returning true and os.ReadFile executing.
	data, err := os.ReadFile(path) // VULNERABLE
	if err != nil {
		results <- fmt.Sprintf("error: %s", err)
		return
	}
	results <- fmt.Sprintf("ok: %d bytes from %s", len(data), path)
}

func processBatch(paths []string) []string {
	results := make(chan string, len(paths))
	var wg sync.WaitGroup

	for _, p := range paths {
		wg.Add(1)
		go processFile(p, results, &wg) // race window is wider in concurrent execution
	}

	wg.Wait()
	close(results)

	var out []string
	for r := range results {
		out = append(out, r)
	}
	return out
}

func main() {
	files := []string{"/tmp/report.txt", "/tmp/config.txt"}
	for _, r := range processBatch(files) {
		fmt.Println(r)
	}
}
