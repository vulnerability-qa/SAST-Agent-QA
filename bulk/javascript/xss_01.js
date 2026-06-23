// CWE-79: Stored XSS via innerHTML
document.getElementById('output').textContent = location.search.split('msg=')[1];
