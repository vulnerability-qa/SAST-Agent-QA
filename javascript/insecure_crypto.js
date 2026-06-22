const crypto = require('crypto');
const bcrypt = require('bcrypt');
const readline = require('readline');

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('Enter password: ', async (password) => {
  // CWE-327: MD5 is cryptographically broken and unsuitable for password hashing.
  const saltRounds = 12;
  const hash = await bcrypt.hash(password, saltRounds);
  console.log('bcrypt:', hash);
  rl.close();
});
