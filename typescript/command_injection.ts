import { exec } from 'child_process';
import * as readline from 'readline';

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter command: ', (cmd: string) => {
  // CWE-78: user input passed to exec without sanitization.
  exec(cmd, (err, stdout) => {
    if (err) { console.error(err); return; }
    console.log(stdout);
  });
  rl.close();
});
