// CWE-79: DOM XSS via document.write
const query = new URLSearchParams(window.location.search);
const div = document.createElement('div');
div.textContent = query.get('name');
document.body.appendChild(div);
