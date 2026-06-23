// CWE-601: Open Redirect
const express = require('express');
const app = express();
app.get('/login', (req, res) => {
  // authenticate...
  const nextUrl = req.query.next;
  if (nextUrl && nextUrl.startsWith('/') && !nextUrl.startsWith('//')) {
    res.redirect(nextUrl);
  } else {
    res.redirect('/');
  }
});
