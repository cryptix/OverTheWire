#include <sys/ptrace.h>
#include <signal.h>
#include <sys/user.h>
#include <sys/types.h>
#include <sys/wait.h>

#include <stdio.h>
#include <unistd.h>
#include <sys/reg.h>

#include <sys/syscall.h>


int main()
{   pid_t child;
    int status = 0, insyscall=0;
    struct user_regs_struct uregs;
    long orig_eax,eax;

    child = fork();
    if(child == 0) {
	//printf("b4fork: euid=%d\n", geteuid());
        ptrace(PTRACE_TRACEME, 0, NULL, NULL);
	//execl("/bin/ls", "ls", NULL);
	// printf("spoofed:%d\n",r);
        execl("/semtex/semtex4", "semtex4", NULL);
    }
    else {
	while(1) {
		wait(&status);

		if(WIFEXITED(status)) {
			printf("We'reee doonnee\n");
			break;
		}

		orig_eax = ptrace(PTRACE_PEEKUSER, child, 4*ORIG_EAX, NULL);
		ptrace(PTRACE_GETREGS,  child, NULL, &uregs);
		printf("syscall #%ld\n", orig_eax);
		if(orig_eax == __NR_geteuid32) {
			if (insyscall == 0) {
				printf("geteuid called");
				insyscall = 1;
			} else {
				//ptrace(PTRACE_GETREGS, child, &uregs);
				printf("geteuid returned with %ld\n", uregs.eax);
				uregs.eax = 6005;
				ptrace(PTRACE_SETREGS, child, NULL, &uregs);
				insyscall = 0;
			}
		}
		ptrace(PTRACE_SYSCALL, child, NULL, NULL);
	}
    }
    return 0;
}
