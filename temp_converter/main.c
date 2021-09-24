
#include <math.h>
#include <stdio.h>
#include <string.h>

const float ftc = 5.0/9.0;
const float ctf = 9.0/5.0;
const float kelz = 238.0;

void usage()
{
    printf("Program convert temperatures \n \
Usage: main ARG DEG, where ARG - convertion type, DEG - temperature \n \
-h   - print this message \n \
-ftc - fahrengeit to celsius \n \
-ctf - celsius to fahrengeit \n");
} 

float CelToFar(float deg)
{
	return ctf*deg + 32.0;
}

float FarToCel(float deg)
{
	return ftc*(deg - 32.0);
}

float CelToKel(float deg)
{
	return deg - kelz;
}

int main(int argc, char *argv[])
{
	float degrees;
	
	if (argc != 3) 
	{
		usage();
		return 0;
	}
	
	if (strcmp(argv[1],"-h") == 0)
		usage();
	
	if (strcmp(argv[1],"-ftc") == 0)
	{
		sscanf(argv[2],"%f",&degrees);
		printf("%.2f \n", FarToCel(degrees));
	}
	
	if (strcmp(argv[1],"-ctf") == 0)
	{
		sscanf(argv[2],"%f",&degrees);
		printf("%.2f \n", CelToFar(degrees));
	}
	
	return 0;
}
