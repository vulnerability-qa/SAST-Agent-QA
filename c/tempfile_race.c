#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main() {
    // CWE-367: tmpnam is vulnerable to TOCTOU race — attacker can replace file between check and open.
    umask(0077);
    char tmpfile_template[] = "/tmp/tempXXXXXX";
    int fd = mkstemp(tmpfile_template);
    if (fd == -1) {
        perror("mkstemp");
        return 1;
    }
    FILE *f = fdopen(fd, "w");
    if (!f) {
        close(fd);
        perror("fdopen");
        return 1;
    }
    fprintf(f, "sensitive data\n");
    fclose(f);
    printf("Wrote to %s\n", tmpfile_template);
    return 0;
}
