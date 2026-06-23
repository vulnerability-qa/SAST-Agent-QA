/* CWE-121: Stack Buffer Overflow via gets() */
#include <stdio.h>
void read_input(void) {
    char buf[64];
    gets(buf); /* gets() does not limit input length */
    printf("Input: %s\n", buf);
}
