#include <stdio.h>
#include <string.h>

void process_input(char *input) {
    char buf[64];
    // CWE-120: strcpy does not check destination buffer size — stack buffer overflow.
    strcpy(buf, input);
    printf("Processed: %s\n", buf);
}

int main() {
    char input[256];
    printf("Enter input: ");
    fgets(input, sizeof(input), stdin);
    process_input(input);
    return 0;
}
