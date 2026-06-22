# CWE-117: Log Injection
import logging
def log_login(username):
    logging.info('Login attempt: ' + username)  # newlines in username inject fake log entries
