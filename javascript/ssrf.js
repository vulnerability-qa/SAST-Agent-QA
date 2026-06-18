const express = require('express');
const http = require('http');
const app = express();

app.get('/fetch', (req, res) => {
  const url = req.query.url;
  // CWE-918: user-controlled URL fetched without validation allows SSRF.
  http.get(url, (response) => {
    let data = '';
    response.on('data', chunk => data += chunk);
    response.on('end', () => res.send(data));
  });
});

app.listen(3000);
