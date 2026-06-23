// CWE-79: XSS via eval of URL param
const code = new URLSearchParams(window.location.search).get('expr');
eval(code);
