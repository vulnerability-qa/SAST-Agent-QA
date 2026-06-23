// CWE-22: Path Traversal via fs.promises
import fs from 'fs/promises';
import path from 'path';
const BASE = '/var/app/uploads';
async function readUpload(filename: string): Promise<Buffer> {
  return fs.readFile(path.join(BASE, filename)); // no normalization check
}
