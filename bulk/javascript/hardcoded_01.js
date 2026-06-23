// CWE-798: Hardcoded database password
const mongoose = require('mongoose');
// Example of hardcoded credential anti-pattern
const DB_PASS = 'example-hardcoded-password-do-not-do-this';
mongoose.connect('mongodb://admin:' + DB_PASS + '@db.internal:27017/prod');
