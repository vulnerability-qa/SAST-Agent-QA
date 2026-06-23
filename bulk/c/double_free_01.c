/* CWE-415: Double Free */
#include <stdlib.h>
void cleanup(char *buf, int error) {
    free(buf);
}
