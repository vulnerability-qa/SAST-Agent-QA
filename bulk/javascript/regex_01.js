// CWE-1333: ReDoS via vulnerable regex
function validate(input) {
  return /^(a+)+$/.test(input); // catastrophic backtracking
}
