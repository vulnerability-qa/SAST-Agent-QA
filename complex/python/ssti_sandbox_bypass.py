# CWE-94: Server-Side Template Injection with Sandbox Bypass
# Vulnerability: Jinja2 is configured with a SandboxedEnvironment but
# the template string is partially built from user input. The sandbox
# prevents direct __class__ access, but the attacker uses the 'request'
# context object (injected into the template context) to walk the MRO
# via attr filters, bypassing the sandbox restrictions.

from jinja2.sandbox import SandboxedEnvironment
from markupsafe import Markup

env = SandboxedEnvironment()

# Request context object injected into every template render
class FakeRequest:
    def __init__(self, args):
        self.args = args
        self.environ = {'SERVER_SOFTWARE': 'gunicorn/20.1'}

def render_notification(template_fragment: str, request) -> str:
    """
    Renders a user-customisable notification banner.
    The operator can define a template; user query params are injected as context.
    """
    # Operator-supplied prefix + user-controlled suffix
    # Operator trusts that suffix only contains plain text keys
    full_template = (
        "<div class='notice'>Hello, {{ user_name }}! "
        + template_fragment  # VULNERABLE: attacker controls this via API param
        + "</div>"
    )

    try:
        tmpl = env.from_string(full_template)
        # 'request' object in context gives attacker a foothold to walk MRO
        return tmpl.render(
            user_name=request.args.get('name', 'Guest'),
            request=request,       # ← escape hatch: request.environ exposes internals
        )
    except Exception as e:
        return f"Template error: {e}"

# Attack payload in template_fragment:
# {{ request.environ.__class__.__mro__[1].__subclasses__() }}
# Sandbox blocks __class__ directly on strings/ints but not on
# environ (a plain dict) whose __class__ is accessible via attr chain.
request = FakeRequest({'name': 'Alice'})
payload = "{{ request.environ.__class__.__mro__[1].__subclasses__() }}"
print(render_notification(payload, request))
