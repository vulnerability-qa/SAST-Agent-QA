// CWE-798: Hardcoded JWT secret
// Anti-pattern: secrets should come from environment variables
const JWT_SECRET = 'example-hardcoded-jwt-secret-never-do-this-in-production';
function sign(payload: object): string {
  const jwt = require('jsonwebtoken');
  return jwt.sign(payload, JWT_SECRET);
}
