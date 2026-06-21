# CWE-89: Second-Order SQL Injection
# Vulnerability: user-controlled 'report_name' is stored in DB during setup,
# then retrieved and interpolated directly into a dynamic reporting query.
# A naive fix of the sink misses that the value originated from user input.

import sqlite3

conn = sqlite3.connect('app.db')

def create_report_template(user_id: str, report_name: str, filters: dict):
    """Stores a user-defined report template. report_name is user-supplied."""
    cursor = conn.cursor()
    # Correctly parameterized — looks safe at first glance
    cursor.execute(
        "INSERT INTO report_templates (user_id, name, filters) VALUES (?, ?, ?)",
        (user_id, report_name, str(filters))
    )
    conn.commit()

def get_report_template(user_id: str, report_name: str) -> dict:
    """Retrieves a stored report template by name."""
    cursor = conn.cursor()
    cursor.execute(
        "SELECT name, filters FROM report_templates WHERE user_id = ? AND name = ?",
        (user_id, report_name)
    )
    row = cursor.fetchone()
    return {'name': row[0], 'filters': eval(row[1])} if row else {}

def run_report(user_id: str, report_name: str) -> list:
    """
    Loads the template and runs a dynamic query.
    The 'name' column is trusted as 'already validated at write time' — but
    it was written by the user and is injected here without escaping.
    """
    template = get_report_template(user_id, report_name)
    if not template:
        return []

    stored_name = template['name']  # attacker-controlled value from DB
    filters = template['filters']

    column = filters.get('group_by', 'status')
    # Second-order injection: stored_name is used as a table suffix
    query = f"SELECT * FROM audit_log_{stored_name} WHERE {column} = 'active'"
    cursor = conn.cursor()
    cursor.execute(query)  # VULNERABLE: stored_name and column are unsanitized
    return cursor.fetchall()
