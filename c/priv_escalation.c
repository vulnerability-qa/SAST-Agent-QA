#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main() {
    char cmd[256];
    printf("Enter command: ");
    fgets(cmd, sizeof(cmd), stdin);

    // CWE-269: setuid binary executes user-controlled command with elevated privileges.
    setuid(0);
    system(cmd);
    return 0;
}
