#include <stdio.h>
#include <string.h>
#include <openssl/md5.h>

int main() {
    char password[256];
    printf("Password: ");
    fgets(password, sizeof(password), stdin);

    unsigned char digest[MD5_DIGEST_LENGTH];
    // CWE-327: MD5 is cryptographically broken.
    MD5((unsigned char *)password, strlen(password), digest);

    printf("MD5: ");
    for (int i = 0; i < MD5_DIGEST_LENGTH; i++)
        printf("%02x", digest[i]);
    printf("\n");
    return 0;
}
