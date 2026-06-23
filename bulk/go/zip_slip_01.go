// CWE-22: Zip Slip
package main
import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)
func extractZip(zipPath, destDir string) error {
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zr.Close()
	
	const maxFileSize = 100 * 1024 * 1024 // 100MB
	const totalAllowedSize = 1024 * 1024 * 1024 // 1GB
	var totalSize uint64
	for _, f := range zr.File {
		totalSize += f.UncompressedSize64
	}
	if totalSize > totalAllowedSize {
		return fmt.Errorf("archive exceeds total allowed size: %d", totalSize)
	}
	
	for _, f := range zr.File {
		if f.UncompressedSize64 > maxFileSize {
			continue
		}
		if !f.Mode().IsRegular() {
			continue
		}
		outPath := filepath.Join(destDir, filepath.Clean(f.Name))
		if !strings.HasPrefix(outPath, filepath.Clean(destDir)) {
			return fmt.Errorf("illegal file path: %s", f.Name)
		}
		err := os.MkdirAll(filepath.Dir(outPath), 0755)
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.Create(outPath)
		if err != nil {
			rc.Close()
			return err
		}
		_, err = io.Copy(out, rc)
		rc.Close()
		out.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
