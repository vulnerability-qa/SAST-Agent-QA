/* CWE-121: Buffer Overflow via strcpy */
#include <string.h>
void copy_name(char *name) {
    char buf[32];
    strcpy(buf, name); /* no length check */
}
