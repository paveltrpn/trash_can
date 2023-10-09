
CC = g++-10
CFLAGS = --std=c++20

SRCS := disasm.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

disasm.o: disasm.cpp
	$(CC) $(CFLAGS) -c disasm.cpp

main: disasm.o
	$(CC) $(OBJS) -o $@ -lstdc++

clean:
	rm -f *.o main 
