// CWE-611: XXE via libxmljs
const libxml = require('libxmljs');
function parseXml(data) {
  return libxml.parseXmlString(data); // external entities enabled by default
}
