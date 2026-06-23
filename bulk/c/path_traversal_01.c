/* CWE-22: Path Traversal via open() */
#include <fcntl.h>
int read_file(const char *filename) {
    char path[256];
    snprintf(path, sizeof(path), "/uploads/%s", filename);
    return open(path, O_RDONLY); /* no normalization */
}
