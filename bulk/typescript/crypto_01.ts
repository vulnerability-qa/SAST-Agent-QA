// CWE-330: Insecure random for session ID
function generateSessionId(): string {
  return Math.random().toString(36).slice(2) + Date.now();
}
