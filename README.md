# SAST-Agent-QA

Vulnerable code samples for testing Legit's Agentic SAST Remediation feature. Each file contains a single intentional vulnerability that should be detected and fixed by the agent.

## Coverage

| Vulnerability | CWE | Python | JS | TS | Java | Go | C# | C |
|---|---|---|---|---|---|---|---|---|
| SQL Injection | CWE-89 | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | |
| Command Injection | CWE-78 | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| XSS | CWE-79 | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | |
| Path Traversal | CWE-22 | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | |
| Insecure Crypto (MD5) | CWE-327 | ✓ | ✓ | | ✓ | ✓ | ✓ | ✓ |
| Insecure Deserialization | CWE-502 | ✓ | | | ✓ | | ✓ | |
| Insecure TLS/Protocol | CWE-295 | ✓ | | | ✓ | ✓ | | |
| CSRF | CWE-352 | ✓ | | | ✓ | | ✓ | |
| SSRF | CWE-918 | ✓ | ✓ | | ✓ | ✓ | | |
| Code / Eval Injection | CWE-94 | ✓ | ✓ | | | | | |
| Prompt Injection | CWE-94 | ✓ | | | | | | |
| Overly Broad Binding | CWE-605 | ✓ | | | | | | |
| World Writable File | CWE-276 | | | | | ✓ | | |
| Insecure Random | CWE-330 | | | | ✓ | | | |
| Privilege Escalation | CWE-269 | | | | ✓ | | | ✓ |
| Spring CSRF Disabled | CWE-352 | | | | ✓ | | | |
| Double Dispose | CWE-674 | | | | | | ✓ | |
| Buffer Overflow | CWE-120 | | | | | | | ✓ |
| Format String | CWE-134 | | | | | | | ✓ |
| Tempfile Race (TOCTOU) | CWE-367 | | | | | | | ✓ |

## Structure

```
├── python/
├── javascript/
├── typescript/
├── java/
├── go/
├── csharp/
└── c/
```

## Purpose

This repository is used to:
- Generate SAST findings across all supported vulnerability types and languages
- Test the end-to-end agentic remediation flow (detection → fix → PR)
- Verify agent behaviour across different severity levels and rule types
- Test the stale patch / concurrent commit scenario

> **Warning:** All code in this repository is intentionally vulnerable. Do not use in production.
