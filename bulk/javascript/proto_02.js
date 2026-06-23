// CWE-1321: Prototype Pollution via URL params
function fromQuery(qs) {
  const obj = {};
  for (const [k, v] of new URLSearchParams(qs)) {
    obj[k] = v; // ?__proto__[isAdmin]=true
  }
  return obj;
}
