// CWE-347: JWT algorithm not restricted — none algorithm accepted
const jwt = require('jsonwebtoken');
function verifyToken(token) {
  // algorithms array not specified — attacker can forge with alg:none
  return jwt.verify(token, 'secret');
}
