// CWE-330: Math.random used for security token
function generateToken() {
  return Math.random().toString(36).substr(2, 16);
}
