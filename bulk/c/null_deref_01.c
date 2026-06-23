/* CWE-476: NULL Pointer Dereference */
#include <string.h>
void get_len(const char *input) {
    /* if input is NULL, strlen crashes */
    size_t len = strlen(input);
}
