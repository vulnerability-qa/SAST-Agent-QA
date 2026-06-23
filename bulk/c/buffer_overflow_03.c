/* CWE-121: Buffer Overflow via sprintf */
#include <stdio.h>
void format_msg(char *user_input) {
    char msg[64];
    snprintf(msg, sizeof(msg), "Hello, %s!", user_input); /* no length check */
}
