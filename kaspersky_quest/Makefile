
CC = g++-10
CFLAGS = --std=c++20

SRCS := test.c
OBJS := $(SRCS:.c=.o)

.phony: all clean

all: main

compile: test.o

test.o: test.c
	gcc -c test.c 

main: $(OBJS)
	$(CC) $(OBJS) -o $@ -lstdc++

clean:
	rm -f *.o main
