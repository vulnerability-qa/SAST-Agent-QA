// CWE-502: Unsafe deserialization via node-serialize
const serialize = require('node-serialize');
function loadSession(cookie) {
  return serialize.unserialize(Buffer.from(cookie, 'base64').toString());
}
