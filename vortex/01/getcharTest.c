#include <stdio.h>
#include <stdlib.h>


int main(int argc, char const *argv[])
{
	unsigned char buf[512];
	unsigned char *ptr = buf + (sizeof(buf)/2);
	unsigned int x;
	int i;

	if (argc == 2)
	{
		fprintf(stderr, "Offsetting: %d\n", atoi(argv[1]));
		// ptr -= atoi(argv[1]);
		for (i = 0; i < atoi(argv[1]); ++i)
		{
			ptr--;
		}

	}

	while((x = getchar()) != EOF) {
		ptr[0] = x;
		printf("%d:%c\n", x, ptr[0]);
	}
	printf("All done.\n");
	return 0;
}