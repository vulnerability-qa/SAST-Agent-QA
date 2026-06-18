import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';

const baseDir = '/var/app/files';
const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter filename: ', (filename: string) => {
  // CWE-22: path.join without normalization check allows directory traversal.
  const fullPath = path.join(baseDir, filename);
  console.log(fs.readFileSync(fullPath, 'utf8'));
  rl.close();
});
