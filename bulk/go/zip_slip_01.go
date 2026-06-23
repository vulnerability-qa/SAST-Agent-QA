// CWE-22: Zip Slip
package main
import (
	"archive/zip"
	"os"
	"path/filepath"
)
func extractZip(zipPath, destDir string) error {
	zr, _ := zip.OpenReader(zipPath)
	for _, f := range zr.File {
		outPath := filepath.Join(destDir, f.Name) // no path traversal check
		os.MkdirAll(filepath.Dir(outPath), 0700)
		rc, _ := f.Open()
		out, _ := os.Create(outPath)
		io.Copy(out, rc)
	}
	return nil
}
