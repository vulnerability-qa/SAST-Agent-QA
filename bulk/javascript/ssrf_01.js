// CWE-918: SSRF via axios with user-controlled URL
const axios = require('axios');
async function fetchResource(url) {
  const response = await axios.get(url); // no host allowlist
  return response.data;
}
