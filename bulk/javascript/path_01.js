// CWE-22: Path Traversal via fs.readFile
const fs = require('fs');
function serveFile(req, res) {
  const file = req.query.file;
  fs.readFile('/var/www/static/' + file, (err, data) => res.send(data));
}
