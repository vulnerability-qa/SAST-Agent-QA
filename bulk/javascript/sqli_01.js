// CWE-89: SQL Injection via string concatenation
const mysql = require('mysql');
const conn = mysql.createConnection({host:'localhost',user:'root',password:'',database:'app'});
function getUser(username) {
  conn.query('SELECT * FROM users WHERE username = \'' + username + '\'', (err, results) => {
    console.log(results);
  });
}
