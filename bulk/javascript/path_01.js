// CWE-22: Path Traversal via fs.readFile
const fs = require('fs');
const path = require('path');
function serveFile(req, res) {
  const file = req.query.file;
  if (!file) {
    return res.status(400).send('File parameter is required');
  }
  if (!/^[a-zA-Z0-9_\-\.]+$/.test(file)) {
    return res.status(400).send('Invalid file name');
  }
  const baseDir = '/var/www/static/';
  const fullPath = path.resolve(path.join(baseDir, file));
  if (!fullPath.startsWith(path.resolve(baseDir))) {
    return res.status(403).send('Access denied');
  }
  fs.readFile(fullPath, (err, data) => {
    if (err) {
      return res.status(404).send('File not found');
    }
    res.setHeader('Content-Type', 'text/plain');
    res.send(data);
  });
}
