const readline = require('readline');
const rl = readline.createInterface({ input: process.stdin, output: process.stdout });

rl.question('Enter code to run: ', (code) => {
  // CWE-477: eval is obsolete and dangerous — executes arbitrary code.
  const result = eval(code);
  console.log('Result:', result);
  rl.close();
});
