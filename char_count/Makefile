
CC = g++-10
CFLAGS = --std=c++20

SRCS := char_count.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

char_count.o: char_count.cpp
	$(CC) $(CFLAGS) -c $<

main: char_count.o 
	$(CC) $(OBJS) -lstdc++ -o main 

clean:
	rm -f *.o main
