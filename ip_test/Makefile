
CC = g++-10
CFLAGS = --std=c++20

SRCS := ip_test.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

ip_test.o: ip_test.cpp
	$(CC) $(CFLAGS) -c ip_test.cpp

main: $(OBJS)
	$(CC) $(OBJS) -o $@ -lstdc++

clean:
	rm -f *.o main
