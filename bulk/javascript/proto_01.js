// CWE-1321: Prototype Pollution via merge
function merge(target, source) {
  for (const key of Object.keys(source)) {
    if (typeof source[key] === 'object') {
      target[key] = merge(target[key] || {}, source[key]);
    } else {
      target[key] = source[key]; // __proto__ can be set here
    }
  }
  return target;
}
