#include <sys/syscall.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>

pid_t geteuid(void) {
  printf("Lulz\n");
  return 666;
}