// CWE-79: DOM XSS via document.write
const query = new URLSearchParams(window.location.search);
document.write('<div>' + query.get('name') + '</div>');
