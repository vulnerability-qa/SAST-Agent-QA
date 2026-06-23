// CWE-22: Path Traversal via res.sendFile without sanitization
const express = require('express');
const path = require('path');
const app = express();
app.get('/download', (req, res) => {
  const baseDir = path.resolve('/uploads/');
  const safeName = path.basename(req.query.name);
  const fullPath = path.join(baseDir, safeName);
  const resolvedPath = path.resolve(fullPath);

  if (!resolvedPath.startsWith(baseDir + path.sep) && resolvedPath !== baseDir) {
    return res.status(400).send('Invalid file path');
  }

  res.sendFile(resolvedPath);
});
