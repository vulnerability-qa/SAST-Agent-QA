// CWE-22: Path Traversal via fs.readFile
const fs = require('fs');
const path = require('path');
function serveFile(req, res) {
  const file = req.query.file;
  const baseDir = '/var/www/static/';
  const safeName = path.basename(file);
  const fullPath = path.join(baseDir, safeName);
  const resolvedPath = path.resolve(fullPath);
  if (!resolvedPath.startsWith(path.resolve(baseDir))) {
    return res.status(400).send('Invalid path specified');
  }
  fs.readFile(resolvedPath, (err, data) => { if (err) return res.status(404).send('File not found'); res.send(data); });
}
