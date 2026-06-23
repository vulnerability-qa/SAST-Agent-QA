// CWE-601: Open Redirect
const express = require('express');
const app = express();
app.get('/login', (req, res) => {
  // authenticate...
  res.redirect(req.query.next || '/');
});
