// CWE-918: SSRF via Multi-Part URL Reconstruction
// Vulnerability: the final URL is assembled from four user-controlled query
// params. Each part passes an individual allowlist check, but combining them
// allows bypassing the domain restriction via a redirect chain or
// credentials-in-URL technique. A fix that only validates the base domain
// is bypassed by the path and fragment components.

import https from 'https';
import { URL } from 'url';

const ALLOWED_DOMAINS = ['api.internal.example.com', 'cdn.example.com'];

function validateDomain(host: string): boolean {
    // Validates host in isolation — does not account for composed URL tricks
    return ALLOWED_DOMAINS.some(d => host === d || host.endsWith('.' + d));
}

function buildWebhookUrl(params: Record<string, string>): string {
    const { scheme, host, path, query, fragment } = params;

    // Each part validated individually — looks thorough
    if (scheme && scheme !== 'https') throw new Error('Only HTTPS allowed');
    if (!validateDomain(host)) throw new Error('Domain not allowed');
    if (path && path.includes('..')) throw new Error('Path traversal blocked');

    // VULNERABLE: assembled URL can encode credentials or bypass via:
    // host=api.internal.example.com&path=@169.254.169.254/latest/meta-data/
    // The '@' in path causes parsers to treat 'api.internal.example.com' as
    // credentials and '169.254.169.254' as the actual host.
    const raw = `${scheme || 'https'}://${host}/${path || ''}${query ? '?' + query : ''}${fragment ? '#' + fragment : ''}`;

    // URL constructor normalises but does not reject credential-bearing URLs
    const url = new URL(raw);
    return url.toString();
}

async function sendWebhook(params: Record<string, string>, body: string): Promise<void> {
    const target = buildWebhookUrl(params); // VULNERABLE
    return new Promise((resolve, reject) => {
        const req = https.request(target, { method: 'POST' }, res => {
            res.on('data', () => {});
            res.on('end', resolve);
        });
        req.on('error', reject);
        req.write(body);
        req.end();
    });
}

export { sendWebhook };
