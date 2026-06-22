# CWE-94: Server-Side Template Injection via Jinja2
from jinja2 import Template
def render(user_template, data):
    return Template(user_template).render(**data)
