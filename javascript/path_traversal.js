const fs = require('fs');
const path = require('path');
const readline = require('readline');

const baseDir = '/var/app/files';
const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter filename: ', (filename) => {
  // CWE-22: no path normalization — allows ../../etc/passwd traversal.
  const fullPath = path.join(baseDir, filename);
  console.log(fs.readFileSync(fullPath, 'utf8'));
  rl.close();
});
