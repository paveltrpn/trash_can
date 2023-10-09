
CC = g++-10
CFLAGS = --std=c++20

SRCS := derivative.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

derivative.o: derivative.cpp
	$(CC) $(CFLAGS) -c derivative.cpp

main: derivative.o
	$(CC) $(OBJS) -o $@ -lstdc++ -lplplotcxx

clean:
	rm -f *.o main
