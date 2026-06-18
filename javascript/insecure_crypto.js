const crypto = require('crypto');
const readline = require('readline');

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter password: ', (password) => {
  // CWE-327: MD5 is cryptographically broken and unsuitable for password hashing.
  const hash = crypto.createHash('md5').update(password).digest('hex');
  console.log('MD5:', hash);
  rl.close();
});
