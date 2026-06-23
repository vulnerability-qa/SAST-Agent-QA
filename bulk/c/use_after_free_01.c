/* CWE-416: Use After Free */
#include <stdlib.h>
#include <string.h>
void process(void) {
    char *buf = malloc(64);
    free(buf);
    memset(buf, 0, 64); /* use after free */
}
