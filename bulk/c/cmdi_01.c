/* CWE-78: Command Injection via system() */
#include <stdlib.h>
void run_command(const char *host) {
    char cmd[256];
    snprintf(cmd, sizeof(cmd), "ping -c 1 %s", host);
    system(cmd); /* host not sanitized */
}
