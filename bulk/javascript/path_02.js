// CWE-22: Path Traversal via res.sendFile without sanitization
const express = require('express');
const app = express();
app.get('/download', (req, res) => {
  res.sendFile('/uploads/' + req.query.name);
});
