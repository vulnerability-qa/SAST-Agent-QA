// CWE-352: CSRF - no token validation on state-changing endpoint
const express = require('express');
const app = express();
app.post('/transfer', (req, res) => {
  const { to, amount } = req.body; // no CSRF token check
  transferFunds(to, amount);
  res.json({ ok: true });
});
function transferFunds() {}
