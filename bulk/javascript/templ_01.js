// CWE-94: Server-Side Template Injection via Handlebars
const Handlebars = require('handlebars');
function render(templateStr, data) {
  const template = Handlebars.compile(templateStr); // user-controlled template
  return template(data);
}
