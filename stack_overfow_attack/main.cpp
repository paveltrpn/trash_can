
#include <iostream>
#include <malloc.h>
#include <memory.h>

const size_t BUF_SIZE = 80;

void greetings(const char *msg) {
    char name[BUF_SIZE];

    strcpy(name, msg);
    printf("Hello, %s\n", name);
}

int main(int argc, char *argv[]) {
    greetings(argv[1]);

    return 0;
}