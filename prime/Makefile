
CC = g++-10
CFLAGS = --std=c++20

SRCS := prime.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

prime.o: prime.cpp
	$(CC) $(CFLAGS) -c prime.cpp 

main: $(OBJS)
	$(CC) $(OBJS) -o $@ -lstdc++

clean:
	rm -f *.o main
