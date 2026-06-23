// CWE-327: MD5 used for password hashing
const crypto = require('crypto');
const bcrypt = require('bcrypt');
async function hashPassword(password) {
  const saltRounds = 12;
  return await bcrypt.hash(password, saltRounds);
}
