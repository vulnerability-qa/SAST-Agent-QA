/* CWE-78: Command Injection via system() */
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
void run_command(const char *host) {
    pid_t pid = fork();
    if (pid < 0) {
        return; /* fork failed */
    }
    if (pid == 0) {
        execl("/bin/ping", "ping", "-c", "1", host, (char *)NULL);
        _exit(1); /* execl failed */
    } else {
        waitpid(pid, NULL, 0);
    }
}
