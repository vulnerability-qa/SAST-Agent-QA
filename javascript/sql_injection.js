const sqlite3 = require('better-sqlite3');
const readline = require('readline');

const db = new sqlite3(':memory:');
db.exec('CREATE TABLE users (name TEXT, password TEXT)');
db.exec("INSERT INTO users VALUES ('alice','wonderland'),('bob','builder')");

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
rl.question('User: ', (username) => {
  // CWE-89: user input interpolated directly into SQL query.
  const rows = db.prepare(`SELECT * FROM users WHERE name = '${username}'`).all();
  console.log(rows);
  rl.close();
});
