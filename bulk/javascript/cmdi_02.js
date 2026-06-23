// CWE-78: Command Injection via execSync
const { execSync } = require('child_process');
function convertImage(filename) {
  return execSync('convert uploads/' + filename + ' output.png').toString();
}
