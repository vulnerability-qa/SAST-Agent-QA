// CWE-918: SSRF via fetch with user-supplied URL
async function fetchWebhook(url: string): Promise<string> {
  const res = await fetch(url); // no allowlist validation
  return res.text();
}
