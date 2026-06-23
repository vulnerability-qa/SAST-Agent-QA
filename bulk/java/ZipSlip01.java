// CWE-22: Zip Slip
import java.io.*;
import java.util.zip.*;
public class ZipSlip01 {
    public void extract(ZipFile zip, File destDir) throws IOException {
        zip.entries().asIterator().forEachRemaining(entry -> {
            File out = new File(destDir, entry.getName()); // no path check
            try { new FileOutputStream(out); } catch (IOException e) {}
        });
    }
}
