import { execFile } from 'child_process';
import * as readline from 'readline';

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter command: ', (cmd: string) => {
  // CWE-78: user input passed to exec without sanitization.
  const allowedCommands = ['ls', 'pwd', 'date'];
  const tokens = cmd.trim().split(/\s+/);
  const commandName = tokens[0];
  const args = tokens.slice(1);

  if (!allowedCommands.includes(commandName)) {
    console.error('Command not allowed');
    rl.close();
    return;
  }

  execFile(commandName, args, (err, stdout) => {
    if (err) { console.error(err); return; }
    console.log(stdout);
  });
  rl.close();
});
