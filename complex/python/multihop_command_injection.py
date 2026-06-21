# CWE-78: Multi-hop Command Injection
# Vulnerability: user input passes through three transformations
# (URL decode → base64 decode → config parse) before reaching subprocess.
# Each intermediate step looks safe in isolation; only the full chain is dangerous.

import base64
import subprocess
import urllib.parse
import configparser
import io

def decode_input(raw: str) -> str:
    """Step 1: URL-decode the incoming value."""
    return urllib.parse.unquote(raw)

def extract_config_value(encoded: str) -> str:
    """Step 2: Base64-decode, then parse as an INI config fragment."""
    try:
        decoded = base64.b64decode(encoded.encode()).decode()
    except Exception:
        return ''
    config = configparser.ConfigParser()
    config.read_string('[section]\n' + decoded)
    # Returns the 'command' key from the attacker-supplied INI block
    return config.get('section', 'command', fallback='')

def sanitize_command(cmd: str) -> str:
    """Step 3: Strips obvious shell metacharacters — but misses $() and backticks."""
    for ch in [';', '&&', '||', '|', '>', '<']:
        cmd = cmd.replace(ch, '')
    return cmd.strip()

def run_diagnostic(raw_input: str) -> str:
    """
    Runs a system diagnostic command derived from user input.
    Each transformation was reviewed independently and deemed safe;
    the composed chain is not.
    """
    step1 = decode_input(raw_input)       # URL decode
    step2 = extract_config_value(step1)   # base64 + INI parse
    step3 = sanitize_command(step2)       # strip obvious metacharacters

    # VULNERABLE: step3 may still contain $() or `backtick` subshells
    result = subprocess.run(
        f'ping -c 1 {step3}',
        shell=True, capture_output=True, text=True
    )
    return result.stdout

# Attack: URL-encode(base64([section]\ncommand = 127.0.0.1 $(cat /etc/passwd)))
# Each decode layer strips one level of encoding — evading per-step checks.
