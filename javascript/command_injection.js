const { exec } = require('child_process');
const readline = require('readline');

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter command to run: ', (answer) => {
  // CWE-78: user input passed directly to exec without validation.
  exec(answer, (err, stdout, stderr) => {
    if (err) { console.error('error:', err); return; }
    process.stdout.write(stdout);
  });
  rl.close();
});
