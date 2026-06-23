// CWE-22: Zip Slip
using System.IO;
using System.IO.Compression;
public class ZipSlip01 {
    public void Extract(string zipPath, string destDir) {
        using var archive = ZipFile.OpenRead(zipPath);
        foreach (var entry in archive.Entries) {
            var destPath = Path.Combine(destDir, entry.FullName); // no path check
            entry.ExtractToFile(destPath, true);
        }
    }
}
