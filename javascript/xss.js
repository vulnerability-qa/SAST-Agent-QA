const express = require('express');
const app = express();

app.get('/greet', (req, res) => {
  const name = req.query.name || '';
  // CWE-79: user input reflected in HTML without escaping.
  res.send(`<h1>Hello, ${name}!</h1>`);
});

app.listen(3000);
