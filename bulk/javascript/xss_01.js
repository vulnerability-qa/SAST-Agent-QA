// CWE-79: Stored XSS via innerHTML
document.getElementById('output').innerHTML = location.search.split('msg=')[1];
