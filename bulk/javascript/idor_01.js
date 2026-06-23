// CWE-639: Insecure Direct Object Reference
const express = require('express');
const app = express();
app.get('/documents/:id', async (req, res) => {
  const doc = await getDocument(req.params.id); // no ownership check
  res.json(doc);
});
async function getDocument(id) { return { id }; }
