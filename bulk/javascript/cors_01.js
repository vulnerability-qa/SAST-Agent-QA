// CWE-942: Overly permissive CORS with credentials
const express = require('express');
const cors = require('cors');
const app = express();
app.use(cors({ origin: '*', credentials: true }));
