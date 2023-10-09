
CC = g++-10
CFLAGS = --std=c++20

SRCS := transpose.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

transpose.o: transpose.cpp
	$(CC) $(CFLAGS) -c transpose.cpp

main: $(OBJS)
	$(CC) $(OBJS) -o $@ -lstdc++

clean:
	rm -f *.o main
