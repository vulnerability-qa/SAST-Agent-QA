// CWE-120: Buffer Overflow via memcpy without size validation
// Detection Rule: Use of memcpy/CopyMemory without size validation may cause
// buffer overflow, leading to crashes or code execution.
// The memcpy family of functions require the developer to validate that the
// destination buffer is the same size or larger than the source buffer.

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define BUFFER_SIZE 64

typedef struct {
    char username[BUFFER_SIZE];
    char session_token[BUFFER_SIZE];
    int  privilege_level;
} UserSession;

void populate_session(UserSession *session, const char *username, const char *token) {
    // VULNERABLE: no validation that strlen(username) < BUFFER_SIZE
    // An oversized username overwrites session_token and privilege_level
    size_t username_len = strnlen(username, BUFFER_SIZE);
    size_t copy_len = (username_len < BUFFER_SIZE) ? username_len : (BUFFER_SIZE - 1);
    memcpy(session->username, username, copy_len);
    session->username[copy_len] = '\0';
    size_t token_len = strnlen(token, BUFFER_SIZE);
    size_t token_copy_len = (token_len < BUFFER_SIZE) ? token_len : (BUFFER_SIZE - 1);
    memcpy(session->session_token, token, token_copy_len);
    session->session_token[token_copy_len] = '\0';
}

void process_request(const char *raw_username, const char *raw_token) {
    UserSession session;
    memset(&session, 0, sizeof(session));
    session.privilege_level = 1; // default: unprivileged

    populate_session(&session, raw_username, raw_token);

    if (session.privilege_level > 1) {
        printf("Admin access granted for: %s\n", session.username);
    } else {
        printf("Standard access for: %s\n", session.username);
    }
}

int main(int argc, char *argv[]) {
    if (argc < 3) {
        fprintf(stderr, "Usage: %s <username> <token>\n", argv[0]);
        return 1;
    }
    // argv values come directly from user input with no length check
    process_request(argv[1], argv[2]);
    return 0;
}
