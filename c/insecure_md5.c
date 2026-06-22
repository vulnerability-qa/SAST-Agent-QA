#include <stdio.h>
#include <string.h>
#include <openssl/evp.h>
#include <openssl/rand.h>

int main() {
    char password[256];
    printf("Password: ");
    fgets(password, sizeof(password), stdin);

    unsigned char salt[16];
    unsigned char derived_key[32];
    int iterations = 100000;
    // Generate random salt
    if (RAND_bytes(salt, sizeof(salt)) != 1) {
        fprintf(stderr, "Error generating random salt\n");
        return 1;
    }
    
    // Derive key using PBKDF2-HMAC-SHA256
    if (PKCS5_PBKDF2_HMAC(password, strnlen(password, sizeof(password)),
                          salt, sizeof(salt),
                          iterations,
                          EVP_sha256(),
                          sizeof(derived_key), derived_key) != 1) {
        fprintf(stderr, "Error deriving key\n");
        return 1;
    }

    printf("Salt: ");
    for (int i = 0; i < sizeof(salt); i++)
        printf("%02x", salt[i]);
    printf("\n");
    
    printf("Derived key: ");
    for (int i = 0; i < sizeof(derived_key); i++)
        printf("%02x", derived_key[i]);
    printf("\n");
    return 0;
}
