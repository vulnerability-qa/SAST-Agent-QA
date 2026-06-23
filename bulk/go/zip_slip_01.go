// CWE-22: Zip Slip
package main
import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)
func extractZip(zipPath, destDir string) error {
	const maxDecompressedSize = 100 * 1024 * 1024 // 100 MB
	zr, _ := zip.OpenReader(zipPath)
	for _, f := range zr.File {
		outPath := filepath.Join(destDir, f.Name) // no path traversal check
		os.MkdirAll(filepath.Dir(outPath), 0755)
		rc, _ := f.Open()
		out, _ := os.Create(outPath)
		limitedReader := io.LimitReader(rc, maxDecompressedSize)
		io.Copy(out, limitedReader)
	}
	return nil
}
