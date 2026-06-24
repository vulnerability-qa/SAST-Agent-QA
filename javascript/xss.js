const express = require('express');
const app = express();

const escapeHtml = (str) => {
  return str.replace(/[&<>"']/g, (m) => {
    return {'&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;'}[m];
  });
};

app.get('/greet', (req, res) => {
  const name = req.query.name || '';
  // CWE-79: user input reflected in HTML without escaping.
  res.send(`<h1>Hello, ${escapeHtml(name)}!</h1>`);
});

app.listen(3000);
