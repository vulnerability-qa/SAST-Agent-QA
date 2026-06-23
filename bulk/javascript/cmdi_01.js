// CWE-78: Command Injection via child_process
const { exec } = require('child_process');
function ping(host) {
  exec('ping -c 1 ' + host, (err, stdout) => console.log(stdout));
}
