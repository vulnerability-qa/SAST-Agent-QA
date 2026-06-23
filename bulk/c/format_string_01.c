/* CWE-134: Format String Injection */
#include <stdio.h>
void log_event(const char *user_str) {
    printf(user_str); /* attacker controls format string */
}
