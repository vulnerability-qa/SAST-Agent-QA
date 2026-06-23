// CWE-918: SSRF via node-fetch
const fetch = require('node-fetch');
async function proxy(req, res) {
  const target = req.query.url;
  const data = await fetch(target).then(r => r.text());
  res.send(data);
}
