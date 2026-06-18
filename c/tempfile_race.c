#include <stdio.h>
#include <stdlib.h>

int main() {
    // CWE-367: tmpnam is vulnerable to TOCTOU race — attacker can replace file between check and open.
    char *tmpfile = tmpnam(NULL);
    FILE *f = fopen(tmpfile, "w");
    if (!f) {
        perror("fopen");
        return 1;
    }
    fprintf(f, "sensitive data\n");
    fclose(f);
    printf("Wrote to %s\n", tmpfile);
    return 0;
}
