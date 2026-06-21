#include <stdio.h>

int main() {
    char buf[256];
    printf("Enter message: ");
    fgets(buf, sizeof(buf), stdin);

    // CWE-134: user input used directly as printf format string.
    printf("%s", buf);
    return 0;
}
