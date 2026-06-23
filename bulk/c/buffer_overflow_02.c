/* CWE-121: Buffer Overflow via strcpy */
#include <string.h>
void copy_name(char *name) {
    char buf[32];
    strncpy(buf, name, sizeof(buf) - 1);
    buf[sizeof(buf) - 1] = '\0';
}
